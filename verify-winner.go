package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Tolsi/vrf-lottery/curve"
	. "github.com/Tolsi/vrf-lottery/tools"
	"github.com/btcsuite/btcutil/base58"
	"io/ioutil"
	"os"
	"reflect"
	"sort"
	"strings"
)

func compare(a, b []string) []string {
	for i := len(a) - 1; i >= 0; i-- {
		for _, vD := range b {
			if a[i] == vD {
				a = append(a[:i], a[i+1:]...)
				break
			}
		}
	}
	return a
}

func main() {
	//region Read params

	provableFilename := flag.String("provableFile", "", "Path to file to validate, it should contains 2 lines: json array of the participants, '\\n' line separator and a block signature")
	blockHeight := flag.Uint("blockHeight", 0, "Waves block height, the signature of it will be used to validate the data")
	proofBase58 := flag.String("proof", "", "Proof bytes in Base58 to validate the message")
	publicKeyBase58 := flag.String("publicKey", "", "curve25519 public key in Base58 to validate the message")
	pickN := flag.Uint("pickN", 1, "Number of winners to pick, it should be >= 1")
	printJson := flag.Bool("json", false, "Output JSON, not plain text")

	flag.Parse()

	if *provableFilename == "" {
		flag.Usage()
		os.Exit(1)
	}
	if *blockHeight == 0 {
		flag.Usage()
		os.Exit(1)
	}
	if *proofBase58 == "" {
		flag.Usage()
		os.Exit(1)
	}
	if *publicKeyBase58 == "" {
		flag.Usage()
		os.Exit(1)
	}
	if *pickN < 1 {
		flag.Usage()
		os.Exit(1)
	}

	participantsAndBlockSignature, err := ioutil.ReadFile(*provableFilename)
	PrintErrorAndExit(err)
	s := strings.Split(string(participantsAndBlockSignature), "\n")
	var participants []string
	err = json.Unmarshal([]byte(s[0]), &participants)
	PrintErrorAndExit(err)
	sortedParticipants := make([]string, len(participants))
	copy(sortedParticipants, participants)
	sort.Strings(sortedParticipants)
	blockSignature := s[1]
	if !reflect.DeepEqual(participants, sortedParticipants) {
		fmt.Printf("Participants list should be sorted \n")
		os.Exit(1)
	}

	proofBytes := base58.Decode(*proofBase58)

	pkb := curve.PublicKey(base58.Decode(*publicKeyBase58))
	PrintErrorAndExit(err)

	//endregion

	//region Create proofs

	loadedBlockSignature, err := GetBlockSignature(uint(*blockHeight))
	PrintErrorAndExit(err)

	if loadedBlockSignature != blockSignature {
		fmt.Printf("Provable file contains different block signature at height %d: expected '%s', got '%s'\n", *blockHeight, loadedBlockSignature, blockSignature)
		os.Exit(1)
	}

	vrfBytes, err := pkb.VerifyVrfSignature(participantsAndBlockSignature, proofBytes)
	if err != nil {
		fmt.Printf("Proof verification was failed\n")
		os.Exit(1)
	}

	//endregion

	//region Result output

	winners := PickUniquePseudorandomParticipants(vrfBytes[:], *pickN, participants)

	if *printJson {
		type OutputWinners struct {
			Winners []string `json:"winners"`
			Proof   string   `json:"proof"`
			Vrf     string   `json:"vrf"`
			Message string   `json:"message"`
		}

		err = json.NewEncoder(os.Stdout).Encode(OutputWinners{winners, *proofBase58, base58.Encode(vrfBytes), string(participantsAndBlockSignature)})
		PrintErrorAndExit(err)
	} else {
		fmt.Printf("message: %s\n", string(participantsAndBlockSignature))
		fmt.Printf("proof (base58): %s\n", *proofBase58)
		fmt.Printf("vrf bytes (base58): %s\n", base58.Encode(vrfBytes))
		fmt.Printf("winners are participants: %v\n", winners)
	}

	//endregion
}
