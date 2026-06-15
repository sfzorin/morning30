package content

// detailsTR is the Turkish rich content for every exercise (current library IDs).
var detailsTR = map[string]Detail{
	// ---- Isınma ----
	"W01": {
		Desc:    "Boyun, omuz ve üst sırtı hazırlayan basit bir omuz ısınması.",
		HowTo:   []string{"Dik durun, kollar gevşek.", "Omuzları yavaşça geriye çevirin.", "Sonra yavaşça öne.", "Hareket akıcı ve kontrollü olsun."},
		Correct: []string{"Boynunuzu gevşek bırakın.", "Omuzlarınızı akıcı şekilde çevirin.", "Kollarınızı serbest tutun.", "Gövdenizi dik tutun."},
		Wrong:   []string{"Omuzlarınızı sertçe yukarı çekmeyin.", "Çok hızlı hareket etmeyin.", "Boynunuzu kasmayın.", "Üst sırtınızı kamburlaştırmayın."},
	},
	"W02": {
		Desc:    "Deltoidleri ve omuz sabitleyicilerini çalıştıran omuz ve üst sırt ısınması.",
		HowTo:   []string{"Dik durun ve kolları omuz hizasına kaldırın.", "Öne küçük daireler çizin.", "Sonra geriye küçük daireler.", "Daireler küçük ve kontrollü olsun."},
		Correct: []string{"Kollarınızı omuz hizasında tutun.", "Daireleri küçük çizin.", "Omuzlarınızı aşağıda tutun.", "Boynunuzu gevşek bırakın."},
		Wrong:   []string{"Büyük savurmalar yapmayın.", "Omuzlarınızı kulaklara kaldırmayın.", "Belinizi çukurlaştırmayın.", "Nefesinizi tutmayın."},
	},
	"W03": {
		Desc:    "Gövde ve üst omurga için yumuşak bir rotasyon çalışması.",
		HowTo:   []string{"Ayaklar kalça genişliğinde.", "Kalça çoğunlukla öne baksın.", "Üst gövdeyi bir yana çevirin.", "Ortaya dönün ve diğer yana çevirin."},
		Correct: []string{"Ayaklarınızı sabit tutun.", "Hareketi gövdenizden başlatın.", "Akıcı şekilde dönün.", "Dizlerinizi burmayın."},
		Wrong:   []string{"Dizlerinizi gövdenizle birlikte çevirmeyin.", "Çok hızlı dönmeyin.", "Geriye yaslanmayın.", "Hareket aralığını zorlamayın."},
	},
	"W04": {
		Desc:    "Kalça, arka bacak ve bel kontrolü için bir hareket çalışması.",
		HowTo:   []string{"Dik durun, ayaklar kalça genişliğinde.", "Dizleri hafif yumuşatın.", "Kalçayı geriye itin.", "Gövdeyi düz sırtla eğin.", "Kalçayı öne iterek doğrulun."},
		Correct: []string{"Sırtınızı düz tutun.", "Hareketi kalçanızdan başlatın.", "Dizlerinizi hafif yumuşak tutun, çok bükmeyin.", "Boynunuzu omurga çizgisinde tutun."},
		Wrong:   []string{"Sırtınızı yuvarlamayın.", "Hareketi çömelmeye çevirmeyin.", "Çok aşağı uzanmayın.", "Başınızı çok kaldırmayın."},
	},
	"W05": {
		Desc:    "Güvenli, sığ bir aralıkta kontrollü bacak ısınması.",
		HowTo:   []string{"Ayaklar kalçadan biraz geniş.", "Kalçayı hafif geriye itin.", "Dizleri yalnızca rahat bir derinliğe bükün.", "Yavaşça doğrulun."},
		Correct: []string{"Dizlerinizi ayak uçlarınızla aynı yönde tutun.", "Topuklarınızı yerde tutun.", "Sırtınızı uzun tutun.", "Derinliği ölçülü tutun."},
		Wrong:   []string{"Çok derine inmeyin.", "Dizlerinizi içe düşürmeyin.", "Topuklarınızı kaldırmayın.", "Hızla çökmeyin."},
	},
	"W06": {
		Desc:    "Omuzlar, bilekler ve kor için bir ısınma.",
		HowTo:   []string{"Yüksek plank pozisyonunda başlayın.", "Eller omuzların altında.", "Gövdeyi hafifçe öne kaydırın.", "Başlangıç pozisyonuna dönün.", "Hareket küçük ve kontrollü."},
		Correct: []string{"Vücudunuzu tek çizgide tutun.", "Omuzlarınızı sabit tutun.", "Karnınızı aktif tutun.", "Akıcı hareket edin."},
		Wrong:   []string{"Kalçanızı düşürmeyin.", "Omuzlarınızı kulaklara itmeyin.", "Çok öne gitmeyin.", "Belinizi çökertmeyin."},
	},

	// ---- Karın / kor ----
	"C01": {
		Desc:    "Karın, omuz, kalça ve gövde için bir kor stabilite egzersizi.",
		HowTo:   []string{"Dirsekleri omuzların altına koyun.", "Ayakları geriye atın.", "Omuzdan topuğa düz çizgi.", "Karın ve kalçayı sıkın.", "Pozisyonu tutun."},
		Correct: []string{"Vücudunuzu tek çizgide tutun.", "Dirseklerinizi omuzlarınızın altında tutun.", "Karnınızı aktif tutun.", "Düzenli nefes alın."},
		Wrong:   []string{"Belinizi çökertmeyin.", "Kalçanızı çok kaldırmayın.", "Öne bakmayın.", "Nefesinizi tutmayın."},
	},
	"C02": {
		Desc:    "Tüm vücut gerilimine odaklı kısa, yoğun bir plank.",
		HowTo:   []string{"Dirsek plankında başlayın.", "Karnı güçlü sıkın.", "Kalçayı sıkın.", "Dirsekleri ayak parmaklarına çektiğinizi düşünün.", "Maksimum kontrolle tutun."},
		Correct: []string{"Tüm vücudunuzu gerin.", "Kalçanızı aktif tutun.", "Karnınızı sıkı tutun.", "Kontrollü nefes alın."},
		Wrong:   []string{"Gevşek bir plank gibi tutmayın.", "Kalçanızı çökertmeyin.", "Nefesinizi tutmayın.", "Çok uzun tutmaya çalışmayın."},
	},
	"C03": {
		Desc:    "Bacakların kontrollü çapraz hareketiyle alt karın egzersizi.",
		HowTo:   []string{"Sırtüstü yatın.", "Beli nazikçe yere bastırın.", "Bacakları kaldırın.", "Bir bacağı diğerinin üzerinden geçirin.", "Çaprazlamayı sürdürün."},
		Correct: []string{"Belinizi yere yakın tutun.", "Bacaklarınızı kontrollü hareket ettirin.", "Boynunuzu gevşek bırakın.", "Akıcı hareket edin."},
		Wrong:   []string{"Belinizi çukurlaştırmayın.", "Bacaklarınızı çok indirmeyin.", "Başınızı öne çekmeyin.", "Çok hızlı hareket etmeyin."},
	},
	"C04": {
		Desc:    "Kısa, almaşık bacak vuruşlarıyla alt karın egzersizi.",
		HowTo:   []string{"Sırtüstü yatın.", "Beli yere bastırın.", "Bacakları kaldırın.", "Bacakları kısa, almaşık şekilde inip kaldırın."},
		Correct: []string{"Hareketi küçük tutun.", "Belinizi sabit tutun.", "Bacaklarınızı çok indirmeyin.", "Düzenli nefes alın."},
		Wrong:   []string{"Belinizi çukurlaştırmayın.", "Büyük vuruşlar yapmayın.", "Boynunuzu kasmayın.", "Nefesinizi tutmayın."},
	},
	"C05": {
		Desc:    "Kalçayı kontrollü kaldırmaya odaklı karın egzersizi.",
		HowTo:   []string{"Sırtüstü yatın.", "Dizleri bükün, ayakları kaldırın.", "Karnı kullanarak kalçayı hafif yukarı kıvırın.", "Yavaşça indirin."},
		Correct: []string{"Alt karnınızı çalıştırın.", "Kalçanızı kontrollü kaldırın.", "Boynunuzu gevşek bırakın.", "Yavaşça indirin."},
		Wrong:   []string{"Bacaklarınızı savurmayın.", "Momentum kullanmayın.", "Boynunuzu çekmeyin.", "Kalçanızı hızla düşürmeyin."},
	},
	"C06": {
		Desc:    "Daha iyi kontrol için kısaltılmış kollu statik kor tutuşu.",
		HowTo:   []string{"Sırtüstü yatın.", "Beli yere bastırın.", "Omuzları hafif kaldırın.", "Dizler bükülü ya da bacaklar daha yukarıda.", "Pozisyonu tutun."},
		Correct: []string{"Belinizi yerde tutun.", "Karnınızı sıkı tutun.", "Boynunuzu gevşek bırakın.", "Nefes almayı sürdürün."},
		Wrong:   []string{"Belinizi çukurlaştırmayın.", "Bacaklarınızı çok indirmeyin.", "Çenenizi öne çekmeyin.", "Nefesinizi tutmayın."},
	},
	"C07": {
		Desc:    "Yan karın, omuz ve kalça stabilitesi için yanal kor egzersizi.",
		HowTo:   []string{"Bir yana yatın.", "Dirseği omuz altına koyun.", "Kalçayı kaldırın.", "Vücut tek düz çizgide.", "Tutun."},
		Correct: []string{"Dirseğinizi omzunuzun altında tutun.", "Kalçanızı yukarıda tutun.", "Vücudunuzu uzun tutun.", "Boynunuzu nötr tutun."},
		Wrong:   []string{"Kalçanızı düşürmeyin.", "Omzunuzu kulağınıza kaldırmayın.", "Göğsünüzü öne yuvarlamayın.", "Nefesinizi tutmayın."},
	},
	"C08": {
		Desc:    "Yan karın ve yanal kalça için daha zorlu bir yan plank çeşidi.",
		HowTo:   []string{"Yan plankta başlayın.", "Kalçayı hafif indirin.", "Kalçayı tekrar kaldırın.", "Kontrolle tekrarlayın.", "Taraf değiştirin."},
		Correct: []string{"Küçük ve kontrollü hareket edin.", "Dirseğinizi omzunuzun altında tutun.", "Kalçanızı dikey hareket ettirin.", "Karnınızı sıkı tutun."},
		Wrong:   []string{"Çok aşağı düşmeyin.", "Gövdenizi burmayın.", "Omzunuza çökmeyin.", "Çok hızlı hareket etmeyin."},
	},
	"C10": {
		Desc:    "Dirsek plankı ile yüksek plank arasında geçişli kor ve triceps egzersizi.",
		HowTo:   []string{"Dirsek plankında başlayın.", "Bir eli yere koyup itin.", "Diğer eli yüksek planka kaldırın.", "Tekrar dirseklere inin.", "Önde gelen kolu değiştirin."},
		Correct: []string{"Kalçanızı sabit tutun.", "Karnınızı aktif tutun.", "Kontrollü hareket edin.", "Omuzlarınızı elinizin ya da dirseğinizin üstünde tutun."},
		Wrong:   []string{"Kalçanızı yana sallamayın.", "Belinizi çökertmeyin.", "Çok hızlı hareket etmeyin.", "Omuzlarınızı çökertmeyin."},
	},
	"C11": {
		Desc:    "Kontrollü bir karın egzersizi ve zor bacak kaldırmalarına daha güvenli bir alternatif.",
		HowTo:   []string{"Sırtüstü yatın.", "Bacakları masa pozisyonuna kaldırın.", "Bir topuğu yere indirin.", "Hafifçe dokunup dönün.", "Tarafları değiştirin."},
		Correct: []string{"Belinizi sabit tutun.", "Yavaş hareket edin.", "Topuğunuzu yumuşakça dokundurun.", "Kalçanızı sabit tutun."},
		Wrong:   []string{"Bacağınızı çok hızlı düşürmeyin.", "Belinizi çukurlaştırmayın.", "İki bacağınızı birlikte hareket ettirmeyin.", "Acele etmeyin."},
	},
	"C12": {
		Desc:    "Derin karın ve omurga stabilitesi için bir kor kontrol egzersizi.",
		HowTo:   []string{"Sırtüstü yatın.", "Kolları ve bacakları masa pozisyonuna kaldırın.", "Karşı kol ve bacağı uzatın.", "Beli sabit tutun.", "Dönün ve taraf değiştirin."},
		Correct: []string{"Belinizi yere yakın tutun.", "Yavaş hareket edin.", "Karşı kolunuzu ve bacağınızı birlikte hareket ettirin.", "Gövdenizi sabit tutun."},
		Wrong:   []string{"Belinizi çukurlaştırmayın.", "Çok hızlı hareket etmeyin.", "Bacağınızı çok indirmeyin.", "Kaburgalarınızı dışarı vermeyin."},
	},

	// ---- Şınav / triceps / omuz ----
	"P01": {
		Desc:    "Göğüs, triceps, omuz ve kor için temel bir üst vücut egzersizi.",
		HowTo:   []string{"Yüksek plankta başlayın.", "Eller omuz altında ya da biraz geniş.", "Vücudu kontrolle indirin.", "Vücudu düz tutarak itin."},
		Correct: []string{"Vücudunuzu tek çizgide tutun.", "Karnınızı sıkı tutun.", "Dirseklerinizi yaklaşık 30-45 derecede tutun.", "Kontrollü hareket edin."},
		Wrong:   []string{"Kalçanızı düşürmeyin.", "Kalçanızı çok kaldırmayın.", "Dirseklerinizi yana açmayın.", "Alt noktaya düşmeyin."},
	},
	"P02": {
		Desc:    "Tricepse daha çok yüklenen bir şınav çeşidi.",
		HowTo:   []string{"Yüksek plankta başlayın.", "Elleri normalden yakın koyun.", "Dirsekleri gövdeye yakın tutun.", "Yavaş inip itin."},
		Correct: []string{"Dirseklerinizi gövdenize yakın tutun.", "Vücudunuzu düz tutun.", "Omuzlarınızı kulaklarınızdan uzak tutun.", "Bileklerinizi rahat tutun."},
		Wrong:   []string{"Ellerinizi çok yakın koymayın.", "Dirseklerinizi açmayın.", "Belinizi çökertmeyin.", "Bilek ya da dirsek ağrısına katlanmayın."},
	},
	"P03": {
		Desc:    "Alt noktada kısa duraklamalı kontrollü şınav.",
		HowTo:   []string{"Normal şınav pozisyonunda başlayın.", "Yavaşça inin.", "Altta bir saniye durun.", "Pozisyonu bozmadan itin."},
		Correct: []string{"Kontrollü duraklayın.", "Karnınızı sıkı tutun.", "Göğsünüzü aktif tutun.", "Vücudunuzu düz tutun."},
		Wrong:   []string{"Altta gevşemeyin.", "Kalçanızı düşürmeyin.", "Nefesinizi tutmayın.", "Alt noktadan sıçramayın."},
	},
	"P04": {
		Desc:    "Kontrolü ve gerilim süresini artıran yavaş şınav çeşidi.",
		HowTo:   []string{"Yüksek plankta başlayın.", "Yaklaşık üç saniye inin.", "Altta bir saniye durun.", "Yaklaşık üç saniye itin."},
		Correct: []string{"Yavaş ve düzenli bir tempo tutun.", "Vücudunuzu düz tutun.", "Dirseklerinizi kontrollü bükün.", "Düzenli nefes alın."},
		Wrong:   []string{"Acele etmeyin.", "Altta çökmeyin.", "Kalçanızı düşürmeyin.", "Karnınızın gerilimini kaybetmeyin."},
	},
	"P05": {
		Desc:    "Göğüs, omuz ve koru zorlayan asimetrik bir şınav çeşidi.",
		HowTo:   []string{"Şınav pozisyonunda başlayın.", "Bir eli hafif önde, diğerini hafif arkada koyun.", "Kontrolle inin.", "İtin.", "Diğer tarafta el konumunu değiştirin."},
		Correct: []string{"Gövdenizi düzgün tutun.", "Karnınızı sıkı tutun.", "İki omzunuzu da sabit tutun.", "Kontrollü hareket edin."},
		Wrong:   []string{"Gövdenizi burmayın.", "Bir omzunuzu düşürmeyin.", "Ellerinizi çok geniş açmayın.", "Omzunuz rahatsızsa aralığı zorlamayın."},
	},
	"P06": {
		Desc:    "Omuza odaklı bir şınav çeşidi.",
		HowTo:   []string{"Yüksek plankta başlayın.", "Ayakları biraz yaklaştırıp kalçayı kaldırın.", "Baş kollar arasında.", "Başı yere doğru indirin.", "Tekrar itin."},
		Correct: []string{"Kalçanızı yukarıda tutun.", "Dirseklerinizi kontrollü bükün.", "Omuzlarınızı sabit tutun.", "Boynunuzu nötr tutun."},
		Wrong:   []string{"Normal şınava düşmeyin.", "Başınızı çok öne götürmeyin.", "Dirseklerinizi çok açmayın.", "Omuz ağrısına katlanmayın."},
	},
	"P08": {
		Desc:    "Önkol desteği pozisyonundan yapılan triceps ve omuz egzersizi.",
		HowTo:   []string{"Önkollar üzerinde, vücut uzun başlayın.", "Dirsekler omuz altında ya da biraz önde.", "Eller ve önkollarla itin.", "Dirsekleri hafif açın.", "Kontrolle dönün."},
		Correct: []string{"Karnınızı sıkı tutun.", "Dirseklerinizi akıcı hareket ettirin.", "Omuzlarınızı aşağıda tutun.", "Hareket aralığınızı rahat tutun."},
		Wrong:   []string{"Dirseklerinizi zorlamayın.", "Kalçanızı düşürmeyin.", "Omuzlarınızı kaldırmayın.", "Dirsek ya da omuz ağrısıyla hareket etmeyin."},
	},

	// ---- Sırt / duruş ----
	"B01": {
		Desc:    "Üst sırt, kürek kemikleri ve omurga kasları için duruş odaklı sırt egzersizi.",
		HowTo:   []string{"Yüzüstü yatın.", "Kollar gövde boyunca.", "Baş parmakları hafif dışa çevirin.", "Göğsü hafif kaldırın.", "Kürek kemiklerini geri ve aşağı çekin.", "Tutun."},
		Correct: []string{"Göğsünüzü az kaldırın.", "Boynunuzu nötr tutun.", "Kürek kemiklerinizi geriye ve aşağıya çekin.", "Kalçanızı hafif aktif tutun."},
		Wrong:   []string{"Çok yükseğe kalkmayın.", "Başınızı geriye atmayın.", "Belinizi aşırı çukurlaştırmayın.", "Omuzlarınızı kulaklara kaldırmayın."},
	},
	"B02": {
		Desc:    "Yüzüstü yapılan üst sırt ve kürek kemiği egzersizi.",
		HowTo:   []string{"Yüzüstü yatın.", "Baş aşağı baksın.", "Kolları yanlardan başa doğru geniş yayla götürün.", "Aynı yoldan dönün.", "Hareket kontrollü."},
		Correct: []string{"Boynunuzu gevşek bırakın.", "Kürek kemiklerinizi akıcı hareket ettirin.", "Göğsünüzü az kaldırın ya da nötr tutun.", "Kollarınızı acele etmeden hareket ettirin."},
		Wrong:   []string{"Başınızı kaldırmayın.", "Belinizi çok çukurlaştırmayın.", "Kollarınızı savurmayın.", "Omuzlarınızı kulaklara kaldırmayın."},
	},
	"B03": {
		Desc:    "Kürek kemikleri ve arka omuz için üst sırt egzersizi.",
		HowTo:   []string{"Yüzüstü yatın.", "Dirsekleri W biçiminde bükün.", "Dirsek ve elleri hafif kaldırın.", "Kürekleri sıkın.", "Kontrolle indirin."},
		Correct: []string{"Hareketi küçük tutun.", "İşi kürek kemiklerinizle yapın.", "Boynunuzu nötr tutun.", "Omuzlarınızı kulaklarınızdan uzak tutun."},
		Wrong:   []string{"Çok yükseğe kaldırmayın.", "Başınızı kaldırmayın.", "Momentum kullanmayın.", "Belinizi aşırı çukurlaştırmayın."},
	},
	"B04": {
		Desc:    "Alt trapez ve üst sırt için kürek kemiği egzersizi.",
		HowTo:   []string{"Yüzüstü yatın.", "Kolları çapraz, baş üstünde Y biçiminde uzatın.", "Kolları yerden hafif kaldırın.", "Kontrolle indirin."},
		Correct: []string{"Kollarınızı yalnızca biraz kaldırın.", "Boynunuzu nötr tutun.", "Omuzlarınızı aşağıda tutun.", "Kontrollü hareket edin."},
		Wrong:   []string{"Çok yükseğe kaldırmayın.", "Omuzlarınızı kulaklara kaldırmayın.", "Belinizi çukurlaştırmayın.", "Çok hızlı hareket etmeyin."},
	},
	"B05": {
		Desc:    "Arka omuz ve kürek kontrolü için üst sırt egzersizi.",
		HowTo:   []string{"Yüzüstü yatın.", "Kolları yanlara T biçiminde açın.", "Kolları hafif kaldırın.", "Kürekleri sıkın.", "Yavaşça indirin."},
		Correct: []string{"Kollarınızı omuz hizasında tutun.", "Boynunuzu gevşek bırakın.", "Kürek kemiklerinizi yumuşakça hareket ettirin.", "Az kaldırın."},
		Wrong:   []string{"Kollarınızı yukarı savurmayın.", "Başınızı kaldırmayın.", "Omuzlarınızı kaldırmayın.", "Geniş aralığı zorlamayın."},
	},
	"B06": {
		Desc:    "Ekipmansız çekiş hareketini taklit eden sırt egzersizi.",
		HowTo:   []string{"Yüzüstü yatın, kollar başın üstünde.", "Göğsü hafif kaldırın.", "Dirsekleri kaburgalara doğru çekin.", "Kürekleri sıkın.", "Kolları tekrar öne uzatın."},
		Correct: []string{"Dirseklerinizi kontrollü indirin.", "Kürek kemiklerinizi hafifçe sıkın.", "Boynunuzu nötr tutun.", "Belinizi rahat tutun."},
		Wrong:   []string{"Kollarınızı sarsmayın.", "Çok yükseğe kalkmayın.", "Başınızı geriye atmayın.", "Omuzlarınızı kulaklara kaldırmayın."},
	},
	"B07": {
		Desc:    "Yüzüstü yapılan sırt ve koordinasyon egzersizi.",
		HowTo:   []string{"Yüzüstü yatın.", "Göğsü hafif kaldırın.", "Karşı kol ve bacağı yüzme gibi hareket ettirin.", "Hareket küçük ve düzenli."},
		Correct: []string{"Boynunuzu nötr tutun.", "Kontrollü hareket edin.", "Belinizi rahat tutun.", "Nefes almayı sürdürün."},
		Wrong:   []string{"Çok hızlı hareket etmeyin.", "Çok yükseğe kalkmayın.", "Başınızı geriye atmayın.", "Belinizi zorlamayın."},
	},
	"B08": {
		Desc:    "Kalça, arka bacak, omuz ve sırt için arka zincir egzersizi.",
		HowTo:   []string{"Bacaklar uzanmış şekilde yere oturun.", "Elleri kalçanın arkasına koyun.", "El ve topuklarla itin.", "Kalçayı kaldırın.", "Vücudu uzun bir çizgide tutun."},
		Correct: []string{"Göğsünüzü açık tutun.", "Kalçanızı yukarıda tutun.", "Omuzlarınızı sabit tutun.", "Boynunuzu nötr tutun."},
		Wrong:   []string{"Kalçanızı düşürmeyin.", "Omuzlarınızı kaldırmayın.", "Başınızı geriye atmayın.", "Bileklerinizi zorlamayın."},
	},
	"B09": {
		Desc:    "Omurga kasları ve duruş için küçük aralıklı sırt egzersizi.",
		HowTo:   []string{"Yüzüstü yatın.", "Eller gövde yanında ya da göğse yakın.", "Göğsü hafif kaldırın.", "Küçük, kontrollü ritmik hareketler yapın.", "Boyun nötr."},
		Correct: []string{"Az kaldırın.", "Akıcı bir ritim tutun.", "Boynunuzu uzun tutun.", "Belinizi kontrollü tutun."},
		Wrong:   []string{"Çok yükseğe çıkmayın.", "Momentum kullanmayın.", "Öne bakmayın.", "Belinizde keskin ağrıya katlanmayın."},
	},
	"B10": {
		Desc:    "Küçük kobra kalkışını W biçimli kol çekişiyle birleştiren daha zorlu sırt egzersizi.",
		HowTo:   []string{"Yüzüstü, kollar önde yatın.", "Göğsü hafif kaldırın.", "Dirsekleri aşağı ve geriye W biçiminde çekin.", "Kürekleri sıkın.", "Tekrar öne uzanıp kontrolle indirin."},
		Correct: []string{"Göğsünüzü az kaldırın.", "W çekişini akıcı yapın.", "Kürek kemiklerinizi geriye ve aşağıya çekin.", "Boynunuzu uzun tutun."},
		Wrong:   []string{"Belinizi aşırı çukurlaştırmayın.", "Boynunuzla çekmeyin.", "Çok hızlı hareket etmeyin.", "Omuzlarınızı kaldırmayın."},
	},

	// ---- Kollar / biceps ----
	"A01": {
		Desc:    "Bir elin diğerine direnç uyguladığı ekipmansız biceps egzersizi.",
		HowTo:   []string{"Dik durun ya da oturun.", "Bir dirseği curl gibi bükün.", "Diğer elle önkola aşağı bastırın.", "Dirence karşı yavaşça yukarı kaldırın.", "Taraf değiştirin."},
		Correct: []string{"Direnci sabit tutun.", "Yavaş hareket edin.", "Dirseğinizi gövdenize yakın tutun.", "Omzunuzu aşağıda tutun."},
		Wrong:   []string{"Gerçek direnç olmadan hareket etmeyin.", "Kolunuzu sarsmayın.", "Omzunuzu kaldırmayın.", "Nefesinizi tutmayın."},
	},
	"A02": {
		Desc:    "Kendine dirençle statik biceps tutuşu.",
		HowTo:   []string{"Dirseği yaklaşık 90 dereceye bükün.", "Diğer elle aşağı bastırın.", "Çalışan kol hareketsiz direnir.", "Gerilimi tutun.", "Taraf değiştirin."},
		Correct: []string{"Dirsek açınızı sabit tutun.", "Gerilimi sabit tutun.", "Omuzlarınızı gevşek bırakın.", "Nefes almayı sürdürün."},
		Wrong:   []string{"Kısa sarsıntılarla bastırmayın.", "Gövdenizi burmayın.", "Omzunuzu kaldırmayın.", "Nefesinizi tutmayın."},
	},
	"A03": {
		Desc:    "Ellerin birbirine karşı kullanıldığı izometrik kol ve önkol egzersizi.",
		HowTo:   []string{"Parmakları ya da elleri kenetleyin.", "Dirsekler hafif bükülü.", "Elleri birbirinden uzağa çekin.", "Sabit gerilim tutun.", "Yavaşça gevşeyin."},
		Correct: []string{"Gerilimi kontrollü tutun.", "Omuzlarınızı aşağıda tutun.", "Bileklerinizi rahat tutun.", "Sakin nefes alın."},
		Wrong:   []string{"Sarsarak çekmeyin.", "Omuzlarınızı kaldırmayın.", "Bileklerinizi aşırı bükmeyin.", "Nefesinizi tutmayın."},
	},
	"A04": {
		Desc:    "Kendine dirençle biceps için yavaş indirme egzersizi.",
		HowTo:   []string{"Bir dirsek bükülü başlayın.", "Diğer elle direnç oluşturun.", "Önkolu yavaşça indirin.", "İndirme boyunca direnci sürdürün.", "Taraf değiştirin."},
		Correct: []string{"Yavaşça indirin.", "Direnci sabit tutun.", "Dirseğinizi gövdenize yakın tutun.", "Omzunuzu gevşek bırakın."},
		Wrong:   []string{"Kolunuzu hızla düşürmeyin.", "Direnci bırakmayın.", "Gövdenizi burmayın.", "Nefesinizi tutmayın."},
	},

	// ---- Bacaklar / kalça ----
	"L01": {
		Desc:    "Bacaklar, kalça ve duruş için kontrollü bir bacak egzersizi.",
		HowTo:   []string{"Ayaklar kalçadan biraz geniş.", "Kalçayı geriye itin.", "Dizleri rahat, sığ bir derinliğe bükün.", "Yavaşça doğrulun."},
		Correct: []string{"Dizlerinizi ayak uçlarınızla aynı yönde tutun.", "Topuklarınızı yerde tutun.", "Sırtınızı uzun tutun.", "Derinliği rahat tutun."},
		Wrong:   []string{"Çok derine inmeyin.", "Dizlerinizi içe düşürmeyin.", "Topuklarınızı kaldırmayın.", "Hızla çökmeyin."},
	},
	"L02": {
		Desc:    "Güvenli aralıkta kısa duraklamalı kontrollü çömelme çeşidi.",
		HowTo:   []string{"Ayaklar sabit durun.", "Rahat, sığ bir çömelmeye inin.", "Kısa süre durun.", "Kontrolle doğrulun."},
		Correct: []string{"Sabit şekilde duraklayın.", "Dizlerinizi ayak uçlarınızla aynı yönde tutun.", "Göğsünüzü açık tutun.", "Topuklarınızı yerde tutun."},
		Wrong:   []string{"Çok derine inmeyin.", "Altta zıplamayın.", "Dizlerinizi içe düşürmeyin.", "Nefesinizi tutmayın."},
	},
	"L03": {
		Desc:    "Kalça, arka bacak ve bel kontrolü için bir hareket çalışması.",
		HowTo:   []string{"Dik durun, ayaklar kalça genişliğinde.", "Dizleri hafif yumuşatın.", "Kalçayı geriye itin.", "Gövdeyi düz sırtla eğin.", "Kalçayı öne iterek doğrulun."},
		Correct: []string{"Sırtınızı düz tutun.", "Hareketi kalçanızdan başlatın.", "Dizlerinizi hafif yumuşak tutun, çok bükmeyin.", "Boynunuzu omurga çizgisinde tutun."},
		Wrong:   []string{"Sırtınızı yuvarlamayın.", "Hareketi çömelmeye çevirmeyin.", "Çok aşağı uzanmayın.", "Başınızı çok kaldırmayın."},
	},
	"L04": {
		Desc:    "Tek bacakla yapılan daha zorlu bir kalça egzersizi.",
		HowTo:   []string{"Sırtüstü yatın.", "Bir dizi bükün, ayak yerde.", "Diğer bacağı uzatın ya da kaldırın.", "Çalışan topuktan itin.", "Kalçayı kaldırıp yavaşça indirin."},
		Correct: []string{"Kalkışı çalışan kalçanızdan yapın.", "Kalçanızı düz tutun.", "Kontrollü hareket edin.", "Belinizi nötr tutun."},
		Wrong:   []string{"Pelvisinizi burmayın.", "Parmak ucunuzdan itmeyin.", "Belinizi çukurlaştırmayın.", "Hızla düşmeyin."},
	},
	"L05": {
		Desc:    "Kalça ve pelvis stabilitesi egzersizi.",
		HowTo:   []string{"Köprü pozisyonunda başlayın.", "Kalçayı yukarıda tutun.", "Bir ayağı yerden hafif kaldırın.", "İndirip taraf değiştirin.", "Pelvis sabit."},
		Correct: []string{"Kalçanızı düz tutun.", "Kalça kaslarınızı aktif tutun.", "Yavaş hareket edin.", "Belinizi rahat tutun."},
		Wrong:   []string{"Kalçanızı düşürmeyin.", "Yana sallanmayın.", "Çok hızlı hareket etmeyin.", "Belinizi çukurlaştırmayın."},
	},
	"L06": {
		Desc:    "Kalça ve arka zinciri güçlendiren bir kalça egzersizi.",
		HowTo:   []string{"Sırtüstü yatın.", "Dizleri bükün, ayaklar yerde.", "Kalçayı kaldırın.", "Üstte durun.", "Yavaşça indirin."},
		Correct: []string{"Topuklarınızdan itin.", "Üstte kalçanızı sıkın.", "Kaburgalarınızı aşağıda tutun.", "Belinizi aşırı açmayın."},
		Wrong:   []string{"Üstte belinizi çukurlaştırmayın.", "Topuklarınızı kaldırmayın.", "Kalçanızı çok hızlı düşürmeyin.", "Ayaklarınızı çok yakın koymayın."},
	},
	"L07": {
		Desc:    "Baldır, ayak bileği ve ayak kontrolü için bir alt bacak egzersizi.",
		HowTo:   []string{"Dik durun, ayaklar kalça genişliğinde.", "Parmak uçlarına yükselin.", "Üstte kısa durun.", "Topukları yavaşça indirin."},
		Correct: []string{"Vücudunuzu dik tutun.", "Akıcı hareket edin.", "Ayak bileklerinizi hizalı tutun.", "Kontrollü indirin."},
		Wrong:   []string{"Topuklarınızı hızla düşürmeyin.", "Öne sallanmayın.", "Ayak bileklerinizi dışa devirmeyin.", "Momentum kullanmayın."},
	},
	"L09": {
		Desc:    "Yan kalça, yan karın ve kalça stabilitesi için yanal zincir egzersizi.",
		HowTo:   []string{"Yan plankta başlayın.", "Vücut düz çizgide.", "Üstteki bacağı hafif kaldırın.", "Kontrolle indirin.", "Tekrarlayıp taraf değiştirin."},
		Correct: []string{"Kalçanızı yukarıda tutun.", "Üst bacağınızı yavaş hareket ettirin.", "Gövdenizi sabit tutun.", "Omzunuzu güçlü tutun."},
		Wrong:   []string{"Kalçanızı düşürmeyin.", "Bacağınızı savurmayın.", "Geriye yuvarlanmayın.", "Omzunuza çökmeyin."},
	},

	// ---- Soğuma ----
	"CD01": {
		Desc:    "Göğüs ve omuz önü için nazik bir esneme.",
		HowTo:   []string{"Sırtüstü yatın.", "Kolları yanlara açın.", "Omuzları gevşetin.", "Göğsün nazikçe açılmasına izin verin.", "Yavaş nefes alın."},
		Correct: []string{"Göğsünüzü yumuşakça açın.", "Omuzlarınızı gevşek bırakın.", "Belinizi rahat tutun.", "Omzunuz ağrımasın."},
		Wrong:   []string{"Kollarınızı zorla bastırmayın.", "Belinizi çukurlaştırmayın.", "Boynunuzu kasmayın.", "Ağrıya rağmen esnemeyin."},
	},
	"CD02": {
		Desc:    "Uyluk arkası için nazik bir esneme.",
		HowTo:   []string{"Sırtüstü yatın.", "Bir bacağı kaldırın.", "Uyluk ya da baldırın arkasından tutun.", "Bacağı nazikçe kendinize çekin.", "Taraf değiştirin."},
		Correct: []string{"Hafifçe esneyin.", "Dizinizi hafif bükülü tutabilirsiniz.", "Belinizi sakin tutun.", "Yavaş nefes alın."},
		Wrong:   []string{"Çok sert çekmeyin.", "Bacağınızı zorla düzleştirmeyin.", "Kalçanızı kaldırmayın.", "Diz arkanızda keskin ağrıya katlanmayın."},
	},
	"CD03": {
		Desc:    "Omurga ve gövde için nazik bir dönüş esnemesi.",
		HowTo:   []string{"Sırtüstü yatın.", "Dizleri bükün.", "Dizleri nazikçe bir yana bırakın.", "Omuzları yerde tutun.", "Ortaya dönüp taraf değiştirin."},
		Correct: []string{"Yavaş hareket edin.", "Omuzlarınızı yerde tutun.", "Gevşek nefes alın.", "Rahatça esneyin."},
		Wrong:   []string{"Dizlerinizi zorla yere bastırmayın.", "Hızlı dönmeyin.", "Omzunuzu kaldırmayın.", "Ağrıya rağmen esnemeyin."},
	},
	"CD04": {
		Desc:    "Nazik bir ön gövde esnemesi ve hafif sırt uzaması.",
		HowTo:   []string{"Yüzüstü yatın.", "Önkolları yere koyun.", "Göğsü nazikçe kaldırın.", "Boynu uzun tutun.", "Yavaş nefes alın."},
		Correct: []string{"Göğsünüzü hafifçe kaldırın.", "Omuzlarınızı kulaklarınızdan uzak tutun.", "Belinizi rahat tutun.", "Boynunuzu nötr tutun."},
		Wrong:   []string{"Belinizi aşırı çukurlaştırmayın.", "Başınızı geriye atmayın.", "Omuzlarınızı kaldırmayın.", "Pozisyonu zorlamayın."},
	},
	"CD05": {
		Desc:    "Seansı bitiren sakin bir nefes egzersizi.",
		HowTo:   []string{"Sırtüstü yatın.", "Omuzları ve çeneyi gevşetin.", "Rahatsa bir eli karnınıza koyun.", "Nazikçe nefes alın.", "Yavaşça verin."},
		Correct: []string{"Sakin nefes alın.", "Omuzlarınızı gevşek bırakın.", "Yüzünüzü yumuşak tutun.", "Vücudunuzu sakinleştirin."},
		Wrong:   []string{"Çok güçlü nefes almayın.", "Nefesinizi tutmayın.", "Boynunuzu kasmayın.", "Belinizi çukurlaştırmayın."},
	},

	// ---- Vlad seti eklemeleri (ısınma / kardiyo / pliometri / hamleler) ----
	"W07": {
		Desc:    "Omurgayı uzatmak için nazik bir tam vücut uzanmasıyla derin nefes alan sakin bir açılış çalışması.",
		HowTo:   []string{"Dik durun, ayaklar kalça genişliğinde.", "Nefes alın ve iki kolu başın üstüne uzatın.", "Omurga boyunca nazikçe uzanın.", "Nefes verin ve kolları indirip omuzları bırakın."},
		Correct: []string{"Yavaş ve derin nefes alın.", "Nefes verirken omuzlarınızı gevşetin.", "Kaburgalarınızı aşağıda tutun, belinizi aşırı çukurlaştırmayın.", "Acele etmeden hareket edin."},
		Wrong:   []string{"Nefesinizi tutmayın.", "Belinizi sertçe çukurlaştırmayın.", "Omuzlarınızı kulaklara kaldırmayın.", "Uzanmayı acele etmeyin."},
	},
	"W08": {
		Desc:    "Omurgayı ve gövdeyi ısıtan nazik bir rotasyon çalışması.",
		HowTo:   []string{"Ayaklar kalça genişliğinde durun.", "Kolların serbestçe sallanmasına izin verin.", "Üst gövdeyi bir yana çevirin.", "Akıcı şekilde diğer yana geçin."},
		Correct: []string{"Hareketi gövdenizden başlatın.", "Kalçanızı çoğunlukla öne bakar tutun.", "Ayaklarınızı sabit tutun.", "Akıcı ve düzenli bir tempo tutun."},
		Wrong:   []string{"Dizlerinizi sertçe burmayın.", "Kollarınızı savurmayın.", "Çok hızlı dönmeyin.", "Nefesinizi tutmayın."},
	},
	"W09": {
		Desc:    "Üst vücudu şınavlara ve sıçrama egzersizlerine hazırlayan dinamik bir omuz ve göğüs ısınması.",
		HowTo:   []string{"Dik durun, ayaklar kalça genişliğinde.", "Kolları kontrollü şekilde öne ve geriye sallayın.", "Kollar geriye giderken göğsün doğal olarak açılmasına izin verin.", "Hareketi akıcı ve gevşek tutun."},
		Correct: []string{"Kollarınızı serbestçe hareket ettirin.", "Omuzlarınızı gevşek bırakın.", "Göğsünüzü nazikçe açın.", "Vücudunuzu dik tutun."},
		Wrong:   []string{"Çok sert sallamayın.", "Omuzlarınızı kulaklara kaldırmayın.", "Belinizi çukurlaştırmayın.", "Hareketi kontrolsüz savurmaya çevirmeyin."},
	},
	"C13": {
		Desc:    "Yan karın ve ön karın için bir rotasyonlu karın egzersizi.",
		HowTo:   []string{"Sırtüstü yatın.", "Dizleri yukarı getirin.", "Elleri hafifçe başın yanına koyun.", "Gövdeyi çevirip bir dirseği karşı dize doğru götürün.", "Kontrollü bir ritimle taraf değiştirin."},
		Correct: []string{"Dönüşü gövdenizden başlatın.", "Belinizi kontrollü tutun.", "Boynunuzu gevşek bırakın.", "Akıcı hareket edin."},
		Wrong:   []string{"Boynunuzdan çekmeyin.", "Yalnızca dirseklerinizi hareket ettirmeyin.", "Tekrarları acele etmeyin.", "Belinizi çukurlaştırmayın."},
	},
	"B11": {
		Desc:    "Karşı kol ve bacağı almaşık, yüzme benzeri bir hareketle kaldıran sırt egzersizi.",
		HowTo:   []string{"Yüzüstü yatın, kollar başın üstünde uzanmış.", "Göğsü ve bacakları hafifçe kaldırın.", "Bir kolu ve karşı bacağı kaldırın.", "İndirip taraf değiştirin, akıcı şekilde almaşık yapın."},
		Correct: []string{"Az kaldırın.", "Boynunuzu omurga çizgisinde tutun.", "Kalça kaslarınızı aktif tutun.", "Düzenli hareket edin."},
		Wrong:   []string{"Belinizi zorlamayın.", "Başınızı kaldırmayın.", "Sarsıntılı hareket etmeyin.", "Nefesinizi tutmayın."},
	},
	"L10": {
		Desc:    "Bacaklar, kalça ve kalça kasları için temel bir vücut ağırlığı çömelmesi.",
		HowTo:   []string{"Ayaklar yaklaşık omuz genişliğinde durun.", "Kalçayı hafif geriye itin.", "Dizleri bükün ve bir çömelmeye inin.", "Göğsü açık tutun.", "Tüm ayağa basarak doğrulun."},
		Correct: []string{"Dizlerinizi ayak uçlarınızla aynı yönde tutun.", "Topuklarınızı yerde tutun.", "Sırtınızı uzun tutun.", "Derinliği kontrollü tutun."},
		Wrong:   []string{"Dizlerinizi içe düşürmeyin.", "Topuklarınızı kaldırmayın.", "Sırtınızı yuvarlamayın.", "Çok hızlı çökmeyin."},
	},
	"L11": {
		Desc:    "Bacaklar, kalça kasları ve denge için bir alt vücut egzersizi.",
		HowTo:   []string{"Dik durun.", "Bir bacağı geriye adımlayın.", "Kontrolle bir hamleye inin.", "Ayakta durmaya dönmek için ön ayaktan itin.", "Taraf değiştirin."},
		Correct: []string{"Ön dizinizi ayak ucunuzla aynı yönde tutun.", "Gövdenizi dik tutun.", "Geri adımı kontrollü atın.", "Ön topuğunuzu yerde tutun."},
		Wrong:   []string{"Ön dizinizi içe düşürmeyin.", "Çok öne eğilmeyin.", "Hızla çökmeyin.", "Arka bacağınızdan çok sert itmeyin."},
	},
	"L12": {
		Desc:    "Kontrol ve bacak gücü geliştirmek için yavaş yapılan bir çömelme.",
		HowTo:   []string{"Ayaklar omuz genişliğinde durun.", "Yaklaşık üç saniyede inin.", "Altta kısa süre durun.", "Kontrolle doğrulun."},
		Correct: []string{"Yavaş bir tempo tutun.", "Dizlerinizi ayak uçlarınızla aynı yönde tutun.", "Topuklarınızı yerde tutun.", "Sırtınızı nötr tutun."},
		Wrong:   []string{"Hızla çökmeyin.", "Alt noktadan sıçramayın.", "Dizlerinizi içe düşürmeyin.", "Nefesinizi tutmayın."},
	},
	"J01": {
		Desc:    "Küçük, hızlı sıçramalarla yapılan hafif bir ayak bileği ve baldır ısınması.",
		HowTo:   []string{"Dik durun, ayaklar yaklaşık kalça genişliğinde.", "Dizleri hafif yumuşak tutun.", "Çoğunlukla ayak bileklerini kullanarak küçük, hızlı sıçramalar yapın.", "Sessizce inip tekrarlayın."},
		Correct: []string{"Alçak sıçrayın.", "Yumuşak ve sessizce inin.", "Dizlerinizi hafif bükülü tutun.", "Vücudunuzu dik tutun."},
		Wrong:   []string{"Çok yükseğe zıplamayın.", "Gürültülü inmeyin.", "Dizlerinizi kilitlemeyin.", "Ayak bileklerinizi içe ya da dışa devirmeyin."},
	},
	"J02": {
		Desc:    "Yüksek diz çekmeye düşük etkili bir ısınma alternatifi.",
		HowTo:   []string{"Dik durun.", "Bir dizi kalça hizasına kaldırın.", "Kontrolle indirin.", "Taraf değiştirip yürüyüşe devam edin.", "Kolları doğal şekilde hareket ettirin."},
		Correct: []string{"Gövdenizi dik tutun.", "Dizlerinizi kontrollü kaldırın.", "Ayaklarınızı yumuşakça indirin.", "Düzenli nefes alın."},
		Wrong:   []string{"Geriye yaslanmayın.", "Bacaklarınızı kontrolsüz savurmayın.", "Ayaklarınızı yere çarpmayın.", "Nefesinizi tutmayın."},
	},
	"J03": {
		Desc:    "Kalp atışını yükselten, omuzları, kalçayı ve bacakları ısıtan tam vücut kardiyo egzersizi.",
		HowTo:   []string{"Dik durun, kollar yanlarda.", "Kolları başın üstüne kaldırırken ayakları yana açın.", "Kolları indirirken ayakları geri bitiştirin.", "Düzenli bir ritimle tekrarlayın."},
		Correct: []string{"Yumuşak inin.", "Dizlerinizi hafif bükülü tutun.", "Kollarınızı akıcı hareket ettirin.", "Ritmik nefes alın."},
		Wrong:   []string{"Sert inmeyin.", "Dizlerinizi kilitlemeyin.", "Omuzlarınızı kulaklara kaldırmayın.", "Çok hızlı hareket edip kontrolü kaybetmeyin."},
	},
	"J04": {
		Desc:    "Yüksek plank pozisyonundan yapılan dinamik bir kor ve kardiyo egzersizi.",
		HowTo:   []string{"Yüksek plankta başlayın.", "Elleri omuzların altında tutun.", "Bir dizi göğse doğru sürün.", "Bacakları ritimle değiştirin.", "Kalçayı sabit tutun."},
		Correct: []string{"Vücudunuzu güçlü bir plankta tutun.", "Karnınızı sıkı tutun.", "Omuzlarınızı ellerinizin üstünde tutun.", "Dizlerinizi kontrollü hareket ettirin."},
		Wrong:   []string{"Kalçanızı yukarı zıplatmayın.", "Belinizi düşürmeyin.", "Ayaklarınızı çok savurmayın.", "Omuzlarınızı çökertmeyin."},
	},
	"J05": {
		Desc:    "Güç, kondisyon ve koordinasyon için pliometrik bir bacak egzersizi.",
		HowTo:   []string{"Ayaklar yaklaşık omuz genişliğinde başlayın.", "Kontrollü bir çömelmeye inin.", "Yukarı zıplayın.", "Dizler hafif bükülü, yumuşak inin.", "Bir sonraki sıçramadan önce pozisyonunuzu toparlayın."},
		Correct: []string{"Sessizce inin.", "Dizlerinizi ayak uçlarınızla aynı yönde tutun.", "Göğsünüzü açık tutun.", "Her sıçramayı kontrollü yapın."},
		Wrong:   []string{"Sert inmeyin.", "Dizlerinizi içe düşürmeyin.", "Çömelmeden zıplamayın.", "Toparlanmadan acele etmeyin."},
	},
	"J06": {
		Desc:    "Bacaklar, kalça, denge ve koordinasyon için yanal bir sıçrama egzersizi.",
		HowTo:   []string{"Tek bacak üzerinde durun.", "Yana doğru diğer bacağa sıçrayın.", "Yumuşak inip dengelenin.", "Yandan yana tekrarlayın.", "Denge için kolları kullanın."},
		Correct: []string{"Sessizce inin.", "Dizinizi ayak ucunuzla aynı yönde tutun.", "Gövdenizi kontrollü tutun.", "Yandan yana akıcı hareket edin."},
		Wrong:   []string{"Dizinizi içe düşürerek inmeyin.", "Çok erken çok uzağa zıplamayın.", "İnişe kendinizi bırakmayın.", "Yere değdiğinizde dizinizi burmayın."},
	},
	"J07": {
		Desc:    "Kalp atışını yükselten, kalçayı, koru ve bacakları aktive eden kardiyo egzersizi.",
		HowTo:   []string{"Dik durun.", "Dizleri yükseğe kaldırarak yerinde koşun.", "Kolları doğal şekilde hareket ettirin.", "Tempoyu hızlı ama kontrollü tutun."},
		Correct: []string{"Dizlerinizi kalça hizasına kaldırın.", "Hafifçe inin.", "Karnınızı aktif tutun.", "Gövdenizi dik tutun."},
		Wrong:   []string{"Geriye yaslanmayın.", "Sert inmeyin.", "Ayaklarınızı yere çarpmayın.", "Tempo arttıkça duruşunuzu kaybetmeyin."},
	},
	"J08": {
		Desc:    "Çömelme, plank ve sıçramayı birleştiren tam vücut kondisyon egzersizi. Bu programda, özellikle şınav atanmadıkça şınavsız çeşidi kullanın.",
		HowTo:   []string{"Dik durun.", "Çömelip elleri yere koyun.", "Ayakları zıplayarak ya da adımlayarak planka atın.", "Ayakları zıplayarak ya da adımlayarak öne getirin.", "Doğrulup küçük bir sıçramayla bitirin."},
		Correct: []string{"Ellerinizi omuzlarınızın altına koyun.", "Güçlü bir plank pozisyonu tutun.", "Ayaklarınızı yumuşakça indirin.", "Kontrollü sıçrayın."},
		Wrong:   []string{"Plankta belinizi düşürmeyin.", "Sert inmeyin.", "Özensiz tekrarlarla acele etmeyin.", "Öne zıplarken dizlerinizi içe düşürmeyin."},
	},
	"J09": {
		Desc:    "Güç, kondisyon ve koordinasyon için ileri seviye pliometrik bir bacak egzersizi.",
		HowTo:   []string{"Hamle pozisyonunda başlayın.", "Yukarı zıplayın.", "Havada bacak değiştirin.", "Karşı tarafta yumuşak bir hamleye inin.", "Bir sonraki tekrardan önce dengenizi toparlayın."},
		Correct: []string{"Yumuşak ve kontrollü inin.", "Ön dizinizi ayak ucunuzla aynı yönde tutun.", "Gövdenizi dik tutun.", "Hareket aralığınızı güvenli tutun."},
		Wrong:   []string{"Sert inmeyin.", "Ön dizinizi içe düşürmeyin.", "Çok derine inmeyin.", "Kontrol edebileceğinizden hızlı hareket etmeyin."},
	},
	"CD07": {
		Desc:    "Sakinleşmek için sırt, kalça ve omuzlara dinlendirici bir esneme.",
		HowTo:   []string{"Diz çöküp kalçayı topuklara doğru geri oturtun.", "Kolları yerde öne uzatın.", "Alnın yere dayanmasına izin verin.", "Yavaşça nefes alıp gevşeyin."},
		Correct: []string{"Kalçanızı topuklarınıza doğru yerleştirin.", "Sırtınızı nazikçe uzatın.", "Omuzlarınızı gevşetin.", "Yavaş nefes alın."},
		Wrong:   []string{"Kalçanızı zorla aşağı bastırmayın.", "Omuzlarınızı kasmayın.", "Nefesinizi tutmayın.", "Dizlerinizi zorlamayın."},
	},
}
