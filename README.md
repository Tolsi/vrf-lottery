## VRF Lottery

This application uses [VRF](https://github.com/coniks-sys/coniks-go/blob/master/crypto/vrf/vrf.go) to confirm the randomness of various online lotteries.
 
The unique provable random index created during the VRF scheme is used to select the winner. 
The winner index is calculated as the remainder of dividing this index by the total number of participants.
Now only tweeter support is made, but you can add any source of participants of the lottery in the future.

# Installation

Download the appropriate file for your platform from the Releases section (windows, mac os linux).

# Scheme of the lottery

To start the lottery, you need to create a public and private keys, select the height of the Waves block in the future, the signature of which will participate in the lottery.
In addition to the conditions for participation in the draw, you need to publicly declare the public key and the height that participants will use in the future for verification.
At the time of summing up the results, the lottery organizer creates a list of participants, adds the block signature at a pre-selected height (saves it as a 'provable file'), selects the winner and creates a 'proof' and a 'vrf' bytes strings).
When announcing the results, the organizer publishes the 'provable file' (list of participants and the block signature), 'vrf' and 'proof'. Using the public key published at the beginning of the lottery, 'provable file', 'vrf' and 'proof', anyone can check the winner, the presence of himself in the list of participants and the impartiality of the organizer.

Since the lottery organizer cannot predict the block signature in the future, he cannot know in advance the pseudo-random value that will be involved in the lottery.
Only the miner of the block in the future may in theory know his signature in advance, but he does not have the private key of the lottery organizer to manipulate it to win the contest, so he does not know how the signature affects the winner.

# Description

The package contains convenient console apps to create keys, load the participants lists, select the winner and check the winner.

Also there're 5 useful native applications:

App `create-keys`
```
Usage: ./create-keys [seed phrase in quotes]
 Output ed25519 private and public keys created from a passed seed phrase.
```

App `pick-winner`
```
Usage: ./pick-winner [participants array file] [waves block height] [proving private key base58 bytes]
 Output winner id (modulo), winner, generate verifiable proof and other usefull info.
```

App `verify-winner`
```
Usage: ./verify-winner [provable file] [vrf base58 bytes] [proof base58 bytes] [proving public key file base58 bytes]
 Output winner id (modulo), winner and verify proof of it.
```

App `retweets-parser`
```
Usage: ./retweets-parser [tweet_id]
 Output retweeters usernames of specified tweet as json array.
```

App `outgoing-waves-tx-recipients`
```
Usage: ./outgoing-waves-tx-recipients [address] [token id] [from ts]
 Output recipients of transfers (only type 4) from this address as json array.
```
 
# Example

```
$ unzip release-linux-*.zip -d vrf-lottery
$ ./create-keys "super mega seed1"
Private key: 265CDdxNGm2XUr6CaQruuNVxqLB18acVqs71ZZiTjqXsj9pKzfB4rh7zCmpLtwSWKyggaXnHFZ54a6GHXp3QQei2
Public key: DuLiPP8KgH11AuiCMACGJ4aXZJypoqFWgRYf4qxYWezJ
$ ./pick-winner participants.txt 1465200 265CDdxNGm2XUr6CaQruuNVxqLB18acVqs71ZZiTjqXsj9pKzfB4rh7zCmpLtwSWKyggaXnHFZ54a6GHXp3QQei2
Provable lottery data was saved to file 'participants_and_1465200_block_signature.txt'
message: ["gerphimum","iaraya","Uyushh","prettyktm","BlankitoOh","TheGobbi","NetCaucusAC","lizan","alexandrelrs","KevCorduroy","mgadda","DuMe","gaberberw","parhamb","dime_looney","mrgrauel","simubucks","01_youtube","TwitterNews","aminechaghal","AijaRuskule1","soumyo_manna","wyattwd","matthewpointon","mankk","ScarthyUK","Aaamra2794","Grachii_xd","dora85997583","Jules_Clarke","Rizal_202122","Hatryst","iWake_Up","konsenkun","RadioInTheUK","teslamint","amadurska","olexhit","explicitfashion","GarethEver","RaihanSelenator","Ana_ipza","nikyve","alecthomas","Speedtao","aliaa_said","AaronMT","TweepsKey","dany_salahudin","Anggreini_L","Dhylsxoxo","ana_0lvr","Evaldnet","canermann","gowdaiam","javiergasilva","Joelysandra","galoeh11","robertvonheeren","JavierJimenez","sunnysiroya","realrisky10","Fyang_23","monchote","ameea31688","kozakmartin","soldierrival","kreyssel","Thelordwar","kkb_03","AM_Adinda","sergioddias","Swifty_Lunatic","xaxuniguiza","MathieuWarnier","iamjeffechano","JulietPatter","setteBIT","OneRainyWish","danielwfiller","quqyqtquqt","prilia_putri12","sarahfriedland","sebasthug","mrjdstokes","rrey","ldijkman1","mrp","AMDeeb90","erikturnquist","bdgfest","Luc_bitch_69","iWarrenMatthews","Ste_ID","taufek_azhar","caiodangelo","fanofkatty","LeosMrs","arifaislam2","gabrielsonio","jonjonweng","CroweSpain","royorbs3","clemtrm","SoyCantuGIIner","MzSassyJ30","XRBOT","fayrouz_moustaf","ameermatani96","RagoZiva","afiyanto1","belizabeth86","steveplunkett","Aluncedo19","CamDvbbs","HD1142200054","nitrotube","Toximityx","khzaw","jsatk","mutia8_indah","twliItter_","9914849662","tBrandii","_Flakaa","WebDevMorgan","SidL_77","jwashbury","marygracechance","ev3927","officialZUL","OpiRoger","aswin_tweet","DominicanaHD","JLBluett","danstudnicky","metti_anggriani","MuellerSimhofer","hiro","callmeabbey","DewiSafitri_30","sum313","Yusrimarcques","sebmysko","maisumrenan","tdh","15june","_hindenbug","pravdadnua","lillerik","EmrysROBERTS","SKSMediaCyprus","asla_waton","afolabitimi","lautlos","DiwaPH","GiantPanda01","ohbargain","ALiffner","Rggprdsng","cbriquet","GiftMitr","buddybuz","mfikril_hakim","UberContent_IN","Mehras_","MoackBELIEBER","hmd","masondixonpro","HugoKessler","Shawneaa31","daemonsy","ruisilva","ostpol","kelleyrobinson","haziimfauzi_","ady_efri","rds_cope","yudhistira_ryan","LoMasDuroRap","MariusNestor","wedharenan","Alfredo_FA14","KeriHenare","ozaaan2","MandeepSinghSEO","SauleyEss","anjas_talluma","lauradyu","threwthatfarawa","CMSLive","morid1n","renan_satiro","luigie_musician","RaminTZ","kevincordova33","sa_srt","eriekx","lheybralston","alfamongi","rainmakertom","FathimaCute","bill220656","Hammadiiii","patrickpang","sbstnhp","Neymar299","naganonopooh","yasarpaslklc","IsraelPatriot","aungkoko134","michaeldardol","CapnLoganWest","LeLeL0","Newsdrom","GenconLt","JJMz12","JohnDomaracki","awaiseaslam","EnergyClimateDE","Ayesha200bcKhan","Minicche","stevenhenman","AMSIOne2Call","iDevSoftware","JazzyDabs","AmritaDeogan","psdesignuk","chrislorquet","gjindal77","abstraktmg","tylrmurphy","jenlxdol","Ariagnanavarro","wintershun","cheeyi","ZHANGY1X1NG","demedtia","VaibhavSharma_","SarahCallaway06","CardRen","1ask_la","nomorepartyinla","iamkgn","scritik","kimberlai12","d7oom1999","monicalopez1313","Neiti_B","gmujaddid","zlizan","CRS_One","bepashahassan47","IIDaddy","patriziagonida","ropowers3","Cheko_Dguez","shanjing","HTheDeep","MrBobbyDigital","GI_Keurmerk","MateusBRib","megalitha_","crsy","ismaelMelendezD","ALCALDEJOS","mul_yadi44","cjtravis","QuinnKTaylor","MostlyVictor","ThufamosiitoRD","087745562868","rpraveens","azucenamark","ac_roca","Pinkiih29","dosankos","alexieasam","ChamoSexii01","hyperlinks2","Ratcicle","_maltem","Naseem864","petererkamp","sanaanm","PMYRS","hejfaisal","fallertiago","zatara214","Mrdoughboy","deps233","aosuzhsbwusnshs","shyamsu09670757","arturgrigor","Manjeet_Paliwal","ResourcefulHR","KevinAMoon","yellow1J","victor__moreira","QuoQlish","vasuri72","RizalFitra_Rif","omegalex1985","MarkJalway","PuteriWe","PetraSlanic","MoAwesomeSauce","ElleVeeZee","butenas_com","qtvr_master","fmcotton","indrafreaks71","viniribeirossa","4GGimbel","Skoua","heartcrazed","chrismessina","Deniseathealthy","tashacorner","CeliaHogan","LaskoLie","nicxmz","grattonboy","buzzaccelerator","mirceap","PutryDestry","ORARiccardo","nunuche5","creativezane","ChiLL_FuNN","matsan41","sharonw","Nutep","2old4dramaok","EsckmoTrent","alistdaily","zZulfan","DavidRadian","LeandroTVmania","nsuaozbausgwnaz","Bonitillo_x2","l2eza","i2workshop","RinaDinapol","darb","sartaj_fawad","SMX815","matthewfarlymn","mdjhaidhosen","geuwmzbsueoxbag","Gay_Germanik","michael_levi","ukcharityevents","RoCurPro2x","smai","jackh71","Livid","__muizz","SWERVEyCHAMp","nickwangler","kalynoh","erkancalp","MichaelTheGeek","metaquest","Jarbouu3","thrivepr","vivapanama90","georg_lange","kumajoez","EJgonzalezc","shinta_yulanda","Younges12021992","JakeSouza19","babar_aadeez","jazzr","PETERNEXTDOOR","Bambang91591296","sustainistmedia","chileandesigner","MichaelBratley","DMaria__","Godknow99939897","yrtytryrytry","COMPUTERWOCHE","kevinforch","ChrisHumphreys_","PuteriSema","ninjafairy89","taquiner9","JoeDiPs","OhMyJet","es_jamely","NinSamolina","dalimoreta12","mintshows","WillsB3","endittriaji","mughni_ef97","Altafkw","SanyaoluMayy","wSuFF","sonyafathinn","weird_one_71126","megoing","Sommer","Dime_Gabriel","ManuelHDInc","jasontm","Mark_Creedon","margotlily","melvynmarrero","DiegoNSalman","Tomatiz","zoomishi","Venuchintam6","Ramoonus","mileyFtDemilol","Intan_Zafira_","bolivar","DanMSargent","rgcottrell","JavierMartinezU","firmanmachda","RDCushing","842961","rockyoctaviano","WEY_LOKO","OLLEMILOKO","jaredgaut","GerriSchmidt1","Andgfaria","ssadowsk","gooochin","sas23ha","elsharmio","redoGH","Jerry00630","pdiamante69","pfreeloader","Pol_McAleese","untitled_av","Team_Brady_","Farideh_Gerami","raymkt","agungputra217","quisye","kyu3","Arbnorswa","Rasztar1","SSimpol","Siingle_02","_RandomVic_","mheederr","miguelfito","SukasaStyle","ididwy","Shoutout400k","TomMorrisEsq","Infomastern","pmcg","Snik3rs_","neahutchins","LarryT63","Lynsay","RJailbreak","svrua","mam0oSh","mako_naka","dvyio","aPhing_Phing","benfloro","ElTes2011","Anthoni3D","Limasbeats","eatotaste","VloggyBuddy","pxixmx","reallyelijahg","carelesstalkk","pam0551","NiceBastard","if__name__","pascal_baratoux","jarh_alamal","abragad","marycbarkley","IsaacIisreal54","rayderds21","9781456739843","Tintin_MASTER","Leaskh","XeekTech","divhealthpsych","OlajideTunde","Craig_Law","JeffreyATW","josandov","contentboosters","madelen_miki","HabibCham","Kuusas","Bahit_oziil","BIG_MOUSE","SARAHGSMILE","alghita_ghifari","FionaWeiland","abdullahijaz100","IoneCoverboy","ollenegro","OMGMediaTV","nicholasnivek","ppl0a","lacarmencitasd","khanumar420","NewElement","carreteromolero","jayman","danliew","babysmurf","hania_azmy","IamJocal","DavidKaneda","MariaDelGreco1","norio_nomura","LeonSYC","MeDiicen_nathy","dfblandfc","Dime_FiiiiiTooo","Believephorn","JuanJArroyo4","Thruth","aradorishirobun","Mr_kamal_bro","3_bittu","jmccray","AMDasanova","Samj0511","itmnrg","MoervYO","GarageSaleAva","TweetSmarter","VitaBrevisBebe","sumanthchoudar1","Amdi1985D","Daniiel08","00amethyst00","gizmosachin","michaelrdk","kgreich","nigzclusive","otaem55","RaeRecum","SassyIdoo","NeoNacho","Wah_You_Want","Hudarohman6","Voceame_Capone","MarcosNet12","AsiwajuShehu","ChellaBrasil","kvox","J_Brackes","chrisdalonzo","m6121","davey","alefrits","Geek_Nurse","manuelacasper2","kuehtarik","MirelaOfydd","rohitnotifies","TomatoWiFi","rodolphedallet","mbrndtgn","italiccreative","bbaekhyun6","adhieravany","mariamuser","fp_rdg","Amilnyleve","Faq505","7even","grnionio","shkodnik83","AdeDj90","sinister_tw","jorgeasolano","errolnardan","purengom","IvonTallo","ChernykhNastasa","wales122","doankmarley","Kim_Rasmussen2","Ruys","MiRO92","esmaelgashuoume","elsananda97","nezhain","NaoEhLeonardo","elkinydicray","alawi5147","glynnpegler","meimanne","hidem511","apfelbasti","kottkrig","mannuelponce","MccBcck","NGUYENTHIMYLIE2","dipaliblaskar","WillianMax","bcoyne","jeanchangster","BeingHuman0009","SlabAds","gamerphoebe","ElFutaiim","MARLA3009","tt_p10","mcleankendraa","kerleniuzmerwez","AlkuwariDr","Lucidreaming123","RojanCetinkaya","netmarketr","MusicWithMJ","nathaly76100259","kost_ru","billchase2edu","yuzuf_Damarzz","writeca","danielrodrigeez","foxyladydi47","Krendelrus","tarasmi","laurentbac","AlbertAdenuga","karvem","BayuBPWA","oncuetelecom","Fowsiyjalaqsan","OxIIID","TwitterIndia","pwscout2","nshsusksbsuskwj","webprolific","jydesign","ronan92","deargem","nya3jp","oseidingba95","AdvocateHafeez","Galuhsadewo","chris27stirling","migueljoseph868","gwilliams782hot","xWindowZ","DimejoseRd","ener","TuDimeMarco","NewParadiseRol","SecurityCrap","jessicanp10","RealHeDi","m_LOjaiN","tuloquiita1","pedroguedesbm","Jack_T11","ilpastafariano","mikewright187","jzetina_","mcapte","lyujunwei_","wilfredphua","pounds_down","abdulghany91","Eddy_HMI","nsox_","li_yewpeng","kabulmuhjib","OneCheck37","numbleroot","tonnydan","vedraniu","jeff_smith","jandet","tuty_utut","apindling242","mipstian","bryandmlee","JMHHACKER","nitolchakma4","nyc_lovers","bidoziad","Bronxstar136","joeyyosa","eikedrescher","k_bujung","polok_jahat","suewang223","evelyn9618","ImWatchingToo","Ed7789","caner58130175","ihternacional","colorfulfigure","iTjensen","TheSMSteam","CarlJeremy11","estrelladelma","nxsueeudbdususi","kankanonmai","billchase2","Maxnl2","melnyk_om","kristw","lubjo81","theegyptiandoll","alepadim","mkonutgan","mrsnormajeancla","dulitharw","PaddyGordon","Isabela69673132","colb_as_ice_","ionull","lecaros","ParksoyeonKr","FlowMelina","maxfriedrich","welles","tom1cho","ElidiahSeruni","Om_Athbii","srta_naranjo_","fadlianggacakep","ohtweet","wentwistle","Enertis","SebFargier","Ju_Vallejo","MarceloAvalos","Wakhe6320","46_alvin","Lightning813","Letcia62410463","madelencyta","williamhayden17","Arif_Widoseno","sofo","shawn_humphrey","ozkatz100","ProphetDLYoung","twit_JN","junpeiwada","wreckseal","DianeFopa","UKGalas","ftvvvv1","Salmon6Peter","GooRetweet","meshal_026","suyash","_risingdigital","matthew_c_perry","Santiag_Vergara","celsocelante","evefavretto","chriskimmelshue","ScrapHeap_ER","najihahhkarim","TiwittrMe","guidomb","yunusPrd","Dj_Newboy","0732917101","syyauqi","arbeitstier","javier_ntobul","Stefanos_Ragos","FaesIlkaMA","Awesome3li","muhammadefendy7","Oldlady12345","xhacker","C_Coolidge","mohuyachowdhur1","mikegapinski","MaribelRecarro"]
djwthTp2Eb2iH85wwmv9peQ5GiShDmAmdJ5bFRTfXigNdapxBrgBqjc6vGfooR48hovgXaBLXZgKR73R2wRVh5e
proof (base58): 2SNkivfTWr3EygcTBWD8QQKLe9x5KxWYCEYskAJjJEso5Rwy7WBoKYZCxALuUQ6YmJ4GAowuP47kNfVpgFMeJZQU3ARQzWw75NXJ3JMdL1anzzBKxYfGuj1NhWi4Vg6kvyxK
vrf bytes (base58): 4bF6ikx8WhQXTpEC57t6gEhSqixio2oUMUzXDs98CwQ2
vrf as number: 24129388030147736487278273252706953108407372648399475657192137631328097176951
modulo: 431
winner is participant #431: ssadowsk
$ ./verify-winner participants_and_1465200_block_signature.txt 1465200 4bF6ikx8WhQXTpEC57t6gEhSqixio2oUMUzXDs98CwQ2 2SNkivfTWr3EygcTBWD8QQKLe9x5KxWYCEYskAJjJEso5Rwy7WBoKYZCxALuUQ6YmJ4GAowuP47kNfVpgFMeJZQU3ARQzWw75NXJ3JMdL1anzzBKxYfGuj1NhWi4Vg6kvyxK DuLiPP8KgH11AuiCMACGJ4aXZJypoqFWgRYf4qxYWezJ
message: ["gerphimum","iaraya","Uyushh","prettyktm","BlankitoOh","TheGobbi","NetCaucusAC","lizan","alexandrelrs","KevCorduroy","mgadda","DuMe","gaberberw","parhamb","dime_looney","mrgrauel","simubucks","01_youtube","TwitterNews","aminechaghal","AijaRuskule1","soumyo_manna","wyattwd","matthewpointon","mankk","ScarthyUK","Aaamra2794","Grachii_xd","dora85997583","Jules_Clarke","Rizal_202122","Hatryst","iWake_Up","konsenkun","RadioInTheUK","teslamint","amadurska","olexhit","explicitfashion","GarethEver","RaihanSelenator","Ana_ipza","nikyve","alecthomas","Speedtao","aliaa_said","AaronMT","TweepsKey","dany_salahudin","Anggreini_L","Dhylsxoxo","ana_0lvr","Evaldnet","canermann","gowdaiam","javiergasilva","Joelysandra","galoeh11","robertvonheeren","JavierJimenez","sunnysiroya","realrisky10","Fyang_23","monchote","ameea31688","kozakmartin","soldierrival","kreyssel","Thelordwar","kkb_03","AM_Adinda","sergioddias","Swifty_Lunatic","xaxuniguiza","MathieuWarnier","iamjeffechano","JulietPatter","setteBIT","OneRainyWish","danielwfiller","quqyqtquqt","prilia_putri12","sarahfriedland","sebasthug","mrjdstokes","rrey","ldijkman1","mrp","AMDeeb90","erikturnquist","bdgfest","Luc_bitch_69","iWarrenMatthews","Ste_ID","taufek_azhar","caiodangelo","fanofkatty","LeosMrs","arifaislam2","gabrielsonio","jonjonweng","CroweSpain","royorbs3","clemtrm","SoyCantuGIIner","MzSassyJ30","XRBOT","fayrouz_moustaf","ameermatani96","RagoZiva","afiyanto1","belizabeth86","steveplunkett","Aluncedo19","CamDvbbs","HD1142200054","nitrotube","Toximityx","khzaw","jsatk","mutia8_indah","twliItter_","9914849662","tBrandii","_Flakaa","WebDevMorgan","SidL_77","jwashbury","marygracechance","ev3927","officialZUL","OpiRoger","aswin_tweet","DominicanaHD","JLBluett","danstudnicky","metti_anggriani","MuellerSimhofer","hiro","callmeabbey","DewiSafitri_30","sum313","Yusrimarcques","sebmysko","maisumrenan","tdh","15june","_hindenbug","pravdadnua","lillerik","EmrysROBERTS","SKSMediaCyprus","asla_waton","afolabitimi","lautlos","DiwaPH","GiantPanda01","ohbargain","ALiffner","Rggprdsng","cbriquet","GiftMitr","buddybuz","mfikril_hakim","UberContent_IN","Mehras_","MoackBELIEBER","hmd","masondixonpro","HugoKessler","Shawneaa31","daemonsy","ruisilva","ostpol","kelleyrobinson","haziimfauzi_","ady_efri","rds_cope","yudhistira_ryan","LoMasDuroRap","MariusNestor","wedharenan","Alfredo_FA14","KeriHenare","ozaaan2","MandeepSinghSEO","SauleyEss","anjas_talluma","lauradyu","threwthatfarawa","CMSLive","morid1n","renan_satiro","luigie_musician","RaminTZ","kevincordova33","sa_srt","eriekx","lheybralston","alfamongi","rainmakertom","FathimaCute","bill220656","Hammadiiii","patrickpang","sbstnhp","Neymar299","naganonopooh","yasarpaslklc","IsraelPatriot","aungkoko134","michaeldardol","CapnLoganWest","LeLeL0","Newsdrom","GenconLt","JJMz12","JohnDomaracki","awaiseaslam","EnergyClimateDE","Ayesha200bcKhan","Minicche","stevenhenman","AMSIOne2Call","iDevSoftware","JazzyDabs","AmritaDeogan","psdesignuk","chrislorquet","gjindal77","abstraktmg","tylrmurphy","jenlxdol","Ariagnanavarro","wintershun","cheeyi","ZHANGY1X1NG","demedtia","VaibhavSharma_","SarahCallaway06","CardRen","1ask_la","nomorepartyinla","iamkgn","scritik","kimberlai12","d7oom1999","monicalopez1313","Neiti_B","gmujaddid","zlizan","CRS_One","bepashahassan47","IIDaddy","patriziagonida","ropowers3","Cheko_Dguez","shanjing","HTheDeep","MrBobbyDigital","GI_Keurmerk","MateusBRib","megalitha_","crsy","ismaelMelendezD","ALCALDEJOS","mul_yadi44","cjtravis","QuinnKTaylor","MostlyVictor","ThufamosiitoRD","087745562868","rpraveens","azucenamark","ac_roca","Pinkiih29","dosankos","alexieasam","ChamoSexii01","hyperlinks2","Ratcicle","_maltem","Naseem864","petererkamp","sanaanm","PMYRS","hejfaisal","fallertiago","zatara214","Mrdoughboy","deps233","aosuzhsbwusnshs","shyamsu09670757","arturgrigor","Manjeet_Paliwal","ResourcefulHR","KevinAMoon","yellow1J","victor__moreira","QuoQlish","vasuri72","RizalFitra_Rif","omegalex1985","MarkJalway","PuteriWe","PetraSlanic","MoAwesomeSauce","ElleVeeZee","butenas_com","qtvr_master","fmcotton","indrafreaks71","viniribeirossa","4GGimbel","Skoua","heartcrazed","chrismessina","Deniseathealthy","tashacorner","CeliaHogan","LaskoLie","nicxmz","grattonboy","buzzaccelerator","mirceap","PutryDestry","ORARiccardo","nunuche5","creativezane","ChiLL_FuNN","matsan41","sharonw","Nutep","2old4dramaok","EsckmoTrent","alistdaily","zZulfan","DavidRadian","LeandroTVmania","nsuaozbausgwnaz","Bonitillo_x2","l2eza","i2workshop","RinaDinapol","darb","sartaj_fawad","SMX815","matthewfarlymn","mdjhaidhosen","geuwmzbsueoxbag","Gay_Germanik","michael_levi","ukcharityevents","RoCurPro2x","smai","jackh71","Livid","__muizz","SWERVEyCHAMp","nickwangler","kalynoh","erkancalp","MichaelTheGeek","metaquest","Jarbouu3","thrivepr","vivapanama90","georg_lange","kumajoez","EJgonzalezc","shinta_yulanda","Younges12021992","JakeSouza19","babar_aadeez","jazzr","PETERNEXTDOOR","Bambang91591296","sustainistmedia","chileandesigner","MichaelBratley","DMaria__","Godknow99939897","yrtytryrytry","COMPUTERWOCHE","kevinforch","ChrisHumphreys_","PuteriSema","ninjafairy89","taquiner9","JoeDiPs","OhMyJet","es_jamely","NinSamolina","dalimoreta12","mintshows","WillsB3","endittriaji","mughni_ef97","Altafkw","SanyaoluMayy","wSuFF","sonyafathinn","weird_one_71126","megoing","Sommer","Dime_Gabriel","ManuelHDInc","jasontm","Mark_Creedon","margotlily","melvynmarrero","DiegoNSalman","Tomatiz","zoomishi","Venuchintam6","Ramoonus","mileyFtDemilol","Intan_Zafira_","bolivar","DanMSargent","rgcottrell","JavierMartinezU","firmanmachda","RDCushing","842961","rockyoctaviano","WEY_LOKO","OLLEMILOKO","jaredgaut","GerriSchmidt1","Andgfaria","ssadowsk","gooochin","sas23ha","elsharmio","redoGH","Jerry00630","pdiamante69","pfreeloader","Pol_McAleese","untitled_av","Team_Brady_","Farideh_Gerami","raymkt","agungputra217","quisye","kyu3","Arbnorswa","Rasztar1","SSimpol","Siingle_02","_RandomVic_","mheederr","miguelfito","SukasaStyle","ididwy","Shoutout400k","TomMorrisEsq","Infomastern","pmcg","Snik3rs_","neahutchins","LarryT63","Lynsay","RJailbreak","svrua","mam0oSh","mako_naka","dvyio","aPhing_Phing","benfloro","ElTes2011","Anthoni3D","Limasbeats","eatotaste","VloggyBuddy","pxixmx","reallyelijahg","carelesstalkk","pam0551","NiceBastard","if__name__","pascal_baratoux","jarh_alamal","abragad","marycbarkley","IsaacIisreal54","rayderds21","9781456739843","Tintin_MASTER","Leaskh","XeekTech","divhealthpsych","OlajideTunde","Craig_Law","JeffreyATW","josandov","contentboosters","madelen_miki","HabibCham","Kuusas","Bahit_oziil","BIG_MOUSE","SARAHGSMILE","alghita_ghifari","FionaWeiland","abdullahijaz100","IoneCoverboy","ollenegro","OMGMediaTV","nicholasnivek","ppl0a","lacarmencitasd","khanumar420","NewElement","carreteromolero","jayman","danliew","babysmurf","hania_azmy","IamJocal","DavidKaneda","MariaDelGreco1","norio_nomura","LeonSYC","MeDiicen_nathy","dfblandfc","Dime_FiiiiiTooo","Believephorn","JuanJArroyo4","Thruth","aradorishirobun","Mr_kamal_bro","3_bittu","jmccray","AMDasanova","Samj0511","itmnrg","MoervYO","GarageSaleAva","TweetSmarter","VitaBrevisBebe","sumanthchoudar1","Amdi1985D","Daniiel08","00amethyst00","gizmosachin","michaelrdk","kgreich","nigzclusive","otaem55","RaeRecum","SassyIdoo","NeoNacho","Wah_You_Want","Hudarohman6","Voceame_Capone","MarcosNet12","AsiwajuShehu","ChellaBrasil","kvox","J_Brackes","chrisdalonzo","m6121","davey","alefrits","Geek_Nurse","manuelacasper2","kuehtarik","MirelaOfydd","rohitnotifies","TomatoWiFi","rodolphedallet","mbrndtgn","italiccreative","bbaekhyun6","adhieravany","mariamuser","fp_rdg","Amilnyleve","Faq505","7even","grnionio","shkodnik83","AdeDj90","sinister_tw","jorgeasolano","errolnardan","purengom","IvonTallo","ChernykhNastasa","wales122","doankmarley","Kim_Rasmussen2","Ruys","MiRO92","esmaelgashuoume","elsananda97","nezhain","NaoEhLeonardo","elkinydicray","alawi5147","glynnpegler","meimanne","hidem511","apfelbasti","kottkrig","mannuelponce","MccBcck","NGUYENTHIMYLIE2","dipaliblaskar","WillianMax","bcoyne","jeanchangster","BeingHuman0009","SlabAds","gamerphoebe","ElFutaiim","MARLA3009","tt_p10","mcleankendraa","kerleniuzmerwez","AlkuwariDr","Lucidreaming123","RojanCetinkaya","netmarketr","MusicWithMJ","nathaly76100259","kost_ru","billchase2edu","yuzuf_Damarzz","writeca","danielrodrigeez","foxyladydi47","Krendelrus","tarasmi","laurentbac","AlbertAdenuga","karvem","BayuBPWA","oncuetelecom","Fowsiyjalaqsan","OxIIID","TwitterIndia","pwscout2","nshsusksbsuskwj","webprolific","jydesign","ronan92","deargem","nya3jp","oseidingba95","AdvocateHafeez","Galuhsadewo","chris27stirling","migueljoseph868","gwilliams782hot","xWindowZ","DimejoseRd","ener","TuDimeMarco","NewParadiseRol","SecurityCrap","jessicanp10","RealHeDi","m_LOjaiN","tuloquiita1","pedroguedesbm","Jack_T11","ilpastafariano","mikewright187","jzetina_","mcapte","lyujunwei_","wilfredphua","pounds_down","abdulghany91","Eddy_HMI","nsox_","li_yewpeng","kabulmuhjib","OneCheck37","numbleroot","tonnydan","vedraniu","jeff_smith","jandet","tuty_utut","apindling242","mipstian","bryandmlee","JMHHACKER","nitolchakma4","nyc_lovers","bidoziad","Bronxstar136","joeyyosa","eikedrescher","k_bujung","polok_jahat","suewang223","evelyn9618","ImWatchingToo","Ed7789","caner58130175","ihternacional","colorfulfigure","iTjensen","TheSMSteam","CarlJeremy11","estrelladelma","nxsueeudbdususi","kankanonmai","billchase2","Maxnl2","melnyk_om","kristw","lubjo81","theegyptiandoll","alepadim","mkonutgan","mrsnormajeancla","dulitharw","PaddyGordon","Isabela69673132","colb_as_ice_","ionull","lecaros","ParksoyeonKr","FlowMelina","maxfriedrich","welles","tom1cho","ElidiahSeruni","Om_Athbii","srta_naranjo_","fadlianggacakep","ohtweet","wentwistle","Enertis","SebFargier","Ju_Vallejo","MarceloAvalos","Wakhe6320","46_alvin","Lightning813","Letcia62410463","madelencyta","williamhayden17","Arif_Widoseno","sofo","shawn_humphrey","ozkatz100","ProphetDLYoung","twit_JN","junpeiwada","wreckseal","DianeFopa","UKGalas","ftvvvv1","Salmon6Peter","GooRetweet","meshal_026","suyash","_risingdigital","matthew_c_perry","Santiag_Vergara","celsocelante","evefavretto","chriskimmelshue","ScrapHeap_ER","najihahhkarim","TiwittrMe","guidomb","yunusPrd","Dj_Newboy","0732917101","syyauqi","arbeitstier","javier_ntobul","Stefanos_Ragos","FaesIlkaMA","Awesome3li","muhammadefendy7","Oldlady12345","xhacker","C_Coolidge","mohuyachowdhur1","mikegapinski","MaribelRecarro"]
djwthTp2Eb2iH85wwmv9peQ5GiShDmAmdJ5bFRTfXigNdapxBrgBqjc6vGfooR48hovgXaBLXZgKR73R2wRVh5e
proof (base58): 2SNkivfTWr3EygcTBWD8QQKLe9x5KxWYCEYskAJjJEso5Rwy7WBoKYZCxALuUQ6YmJ4GAowuP47kNfVpgFMeJZQU3ARQzWw75NXJ3JMdL1anzzBKxYfGuj1NhWi4Vg6kvyxK
vrf bytes (base58): 4bF6ikx8WhQXTpEC57t6gEhSqixio2oUMUzXDs98CwQ2
vrf as number: 24129388030147736487278273252706953108407372648399475657192137631328097176951
modulo: 431
winner is participant #431: ssadowsk
```
