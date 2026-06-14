package content

// detailsTR is the Turkish rich content for every exercise (current library IDs).
var detailsTR = map[string]Detail{
	// ---- Isınma ----
	"W01": {
		Desc:    "Boyun, omuz ve üst sırtı hazırlayan basit bir omuz ısınması.",
		HowTo:   []string{"Dik durun, kollar gevşek.", "Omuzları yavaşça geriye çevirin.", "Sonra yavaşça öne.", "Hareket akıcı ve kontrollü olsun."},
		Correct: []string{"Boyun gevşek.", "Omuzlar akıcı hareket eder.", "Kollar serbest.", "Gövde dik."},
		Wrong:   []string{"Omuzları sertçe yukarı çekmek.", "Çok hızlı hareket etmek.", "Boynu kasmak.", "Üst sırtı kamburlaştırmak."},
	},
	"W02": {
		Desc:    "Deltoidleri ve omuz sabitleyicilerini çalıştıran omuz ve üst sırt ısınması.",
		HowTo:   []string{"Dik durun ve kolları omuz hizasına kaldırın.", "Öne küçük daireler çizin.", "Sonra geriye küçük daireler.", "Daireler küçük ve kontrollü olsun."},
		Correct: []string{"Kollar omuz hizasında.", "Daireler küçük.", "Omuzlar aşağıda.", "Boyun gevşek."},
		Wrong:   []string{"Büyük savurmalar yapmak.", "Omuzları kulaklara kaldırmak.", "Beli çukurlaştırmak.", "Nefesi tutmak."},
	},
	"W03": {
		Desc:    "Gövde ve üst omurga için yumuşak bir rotasyon çalışması.",
		HowTo:   []string{"Ayaklar kalça genişliğinde.", "Kalça çoğunlukla öne baksın.", "Üst gövdeyi bir yana çevirin.", "Ortaya dönün ve diğer yana çevirin."},
		Correct: []string{"Ayaklar sabit.", "Hareket gövdeden gelir.", "Dönüş akıcı.", "Dizler burulmaz."},
		Wrong:   []string{"Dizleri gövdeyle çevirmek.", "Çok hızlı dönmek.", "Geriye yaslanmak.", "Hareket aralığını zorlamak."},
	},
	"W04": {
		Desc:    "Kalça, arka bacak ve bel kontrolü için bir hareket çalışması.",
		HowTo:   []string{"Dik durun, ayaklar kalça genişliğinde.", "Dizleri hafif yumuşatın.", "Kalçayı geriye itin.", "Gövdeyi düz sırtla eğin.", "Kalçayı öne iterek doğrulun."},
		Correct: []string{"Sırt düz.", "Hareket kalçadan başlar.", "Dizler yumuşak, çok bükülmemiş.", "Boyun omurga çizgisini izler."},
		Wrong:   []string{"Sırtı yuvarlamak.", "Çömelmeye çevirmek.", "Çok aşağı uzanmak.", "Başı çok kaldırmak."},
	},
	"W05": {
		Desc:    "Güvenli, sığ bir aralıkta kontrollü bacak ısınması.",
		HowTo:   []string{"Ayaklar kalçadan biraz geniş.", "Kalçayı hafif geriye itin.", "Dizleri yalnızca rahat bir derinliğe bükün.", "Yavaşça doğrulun."},
		Correct: []string{"Dizler parmak yönünde.", "Topuklar yerde.", "Sırt uzun.", "Derinlik ölçülü."},
		Wrong:   []string{"Çok derine inmek.", "Dizleri içe düşürmek.", "Topukları kaldırmak.", "Hızla çökmek."},
	},
	"W06": {
		Desc:    "Omuzlar, bilekler ve kor için bir ısınma.",
		HowTo:   []string{"Yüksek plank pozisyonunda başlayın.", "Eller omuzların altında.", "Gövdeyi hafifçe öne kaydırın.", "Başlangıç pozisyonuna dönün.", "Hareket küçük ve kontrollü."},
		Correct: []string{"Vücut tek çizgide.", "Omuzlar sabit.", "Kor aktif.", "Hareket akıcı."},
		Wrong:   []string{"Kalçayı düşürmek.", "Omuzları kulaklara itmek.", "Çok öne gitmek.", "Beli çökertmek."},
	},

	// ---- Karın / kor ----
	"C01": {
		Desc:    "Karın, omuz, kalça ve gövde için bir kor stabilite egzersizi.",
		HowTo:   []string{"Dirsekleri omuzların altına koyun.", "Ayakları geriye atın.", "Omuzdan topuğa düz çizgi.", "Karın ve kalçayı sıkın.", "Pozisyonu tutun."},
		Correct: []string{"Vücut tek çizgide.", "Dirsekler omuz altında.", "Kor aktif.", "Nefes düzenli."},
		Wrong:   []string{"Beli çökertmek.", "Kalçayı çok kaldırmak.", "Öne bakmak.", "Nefesi tutmak."},
	},
	"C02": {
		Desc:    "Tüm vücut gerilimine odaklı kısa, yoğun bir plank.",
		HowTo:   []string{"Dirsek plankında başlayın.", "Karnı güçlü sıkın.", "Kalçayı sıkın.", "Dirsekleri ayak parmaklarına çektiğinizi düşünün.", "Maksimum kontrolle tutun."},
		Correct: []string{"Tüm vücut gerilimi.", "Kalça aktif.", "Kor sıkı.", "Nefes kontrollü."},
		Wrong:   []string{"Gevşek plank gibi tutmak.", "Kalçayı çökertmek.", "Nefesi tutmak.", "Çok uzun tutmaya çalışmak."},
	},
	"C03": {
		Desc:    "Bacakların kontrollü çapraz hareketiyle alt karın egzersizi.",
		HowTo:   []string{"Sırtüstü yatın.", "Beli nazikçe yere bastırın.", "Bacakları kaldırın.", "Bir bacağı diğerinin üzerinden geçirin.", "Çaprazlamayı sürdürün."},
		Correct: []string{"Bel yere yakın.", "Bacaklar kontrollü.", "Boyun gevşek.", "Hareket akıcı."},
		Wrong:   []string{"Beli çukurlaştırmak.", "Bacakları çok indirmek.", "Başı öne çekmek.", "Çok hızlı hareket etmek."},
	},
	"C04": {
		Desc:    "Kısa, almaşık bacak vuruşlarıyla alt karın egzersizi.",
		HowTo:   []string{"Sırtüstü yatın.", "Beli yere bastırın.", "Bacakları kaldırın.", "Bacakları kısa, almaşık şekilde inip kaldırın."},
		Correct: []string{"Küçük hareket aralığı.", "Bel sabit.", "Bacaklar çok inmez.", "Nefes düzenli."},
		Wrong:   []string{"Beli çukurlaştırmak.", "Büyük vuruşlar yapmak.", "Boynu kasmak.", "Nefesi tutmak."},
	},
	"C05": {
		Desc:    "Kalçayı kontrollü kaldırmaya odaklı karın egzersizi.",
		HowTo:   []string{"Sırtüstü yatın.", "Dizleri bükün, ayakları kaldırın.", "Karnı kullanarak kalçayı hafif yukarı kıvırın.", "Yavaşça indirin."},
		Correct: []string{"Alt karın çalışır.", "Kalça kontrollü kalkar.", "Boyun gevşek.", "İndirme yavaş."},
		Wrong:   []string{"Bacakları savurmak.", "Momentum kullanmak.", "Boynu çekmek.", "Kalçayı hızla düşürmek."},
	},
	"C06": {
		Desc:    "Daha iyi kontrol için kısaltılmış kollu statik kor tutuşu.",
		HowTo:   []string{"Sırtüstü yatın.", "Beli yere bastırın.", "Omuzları hafif kaldırın.", "Dizler bükülü ya da bacaklar daha yukarıda.", "Pozisyonu tutun."},
		Correct: []string{"Bel yerde.", "Karın sıkı.", "Boyun gevşek.", "Nefes sürer."},
		Wrong:   []string{"Beli çukurlaştırmak.", "Bacakları çok indirmek.", "Çeneyi öne çekmek.", "Nefesi tutmak."},
	},
	"C07": {
		Desc:    "Yan karın, omuz ve kalça stabilitesi için yanal kor egzersizi.",
		HowTo:   []string{"Bir yana yatın.", "Dirseği omuz altına koyun.", "Kalçayı kaldırın.", "Vücut tek düz çizgide.", "Tutun."},
		Correct: []string{"Dirsek omuz altında.", "Kalça yukarıda.", "Vücut uzun.", "Boyun nötr."},
		Wrong:   []string{"Kalçayı düşürmek.", "Omzu kulağa kaldırmak.", "Göğsü öne yuvarlamak.", "Nefesi tutmak."},
	},
	"C08": {
		Desc:    "Yan karın ve yanal kalça için daha zorlu bir yan plank çeşidi.",
		HowTo:   []string{"Yan plankta başlayın.", "Kalçayı hafif indirin.", "Kalçayı tekrar kaldırın.", "Kontrolle tekrarlayın.", "Taraf değiştirin."},
		Correct: []string{"Hareket küçük ve kontrollü.", "Dirsek omuz altında.", "Kalça dikey hareket eder.", "Kor sıkı."},
		Wrong:   []string{"Çok aşağı düşmek.", "Gövdeyi burmak.", "Omza çökmek.", "Çok hızlı hareket etmek."},
	},
	"C10": {
		Desc:    "Dirsek plankı ile yüksek plank arasında geçişli kor ve triceps egzersizi.",
		HowTo:   []string{"Dirsek plankında başlayın.", "Bir eli yere koyup itin.", "Diğer eli yüksek planka kaldırın.", "Tekrar dirseklere inin.", "Önde gelen kolu değiştirin."},
		Correct: []string{"Kalça sabit.", "Kor aktif.", "Hareket kontrollü.", "Omuzlar el ya da dirsek üstünde."},
		Wrong:   []string{"Kalçayı yana sallamak.", "Beli çökertmek.", "Çok hızlı hareket etmek.", "Omuzları çökertmek."},
	},
	"C11": {
		Desc:    "Kontrollü bir karın egzersizi ve zor bacak kaldırmalarına daha güvenli bir alternatif.",
		HowTo:   []string{"Sırtüstü yatın.", "Bacakları masa pozisyonuna kaldırın.", "Bir topuğu yere indirin.", "Hafifçe dokunup dönün.", "Tarafları değiştirin."},
		Correct: []string{"Bel sabit.", "Hareket yavaş.", "Topuk yumuşak dokunur.", "Kalça sabit."},
		Wrong:   []string{"Bacağı çok hızlı düşürmek.", "Beli çukurlaştırmak.", "İki bacağı birlikte hareket ettirmek.", "Acele etmek."},
	},
	"C12": {
		Desc:    "Derin karın ve omurga stabilitesi için bir kor kontrol egzersizi.",
		HowTo:   []string{"Sırtüstü yatın.", "Kolları ve bacakları masa pozisyonuna kaldırın.", "Karşı kol ve bacağı uzatın.", "Beli sabit tutun.", "Dönün ve taraf değiştirin."},
		Correct: []string{"Bel yere yakın.", "Hareket yavaş.", "Karşı kol ve bacak birlikte hareket eder.", "Gövde sabit."},
		Wrong:   []string{"Beli çukurlaştırmak.", "Çok hızlı hareket etmek.", "Bacağı çok indirmek.", "Kaburgaları dışarı vermek."},
	},

	// ---- Şınav / triceps / omuz ----
	"P01": {
		Desc:    "Göğüs, triceps, omuz ve kor için temel bir üst vücut egzersizi.",
		HowTo:   []string{"Yüksek plankta başlayın.", "Eller omuz altında ya da biraz geniş.", "Vücudu kontrolle indirin.", "Vücudu düz tutarak itin."},
		Correct: []string{"Vücut tek çizgide.", "Kor sıkı.", "Dirsekler yaklaşık 30–45 derece.", "Hareket kontrollü."},
		Wrong:   []string{"Kalçayı düşürmek.", "Kalçayı çok kaldırmak.", "Dirsekleri yana açmak.", "Alt noktaya düşmek."},
	},
	"P02": {
		Desc:    "Tricepse daha çok yüklenen bir şınav çeşidi.",
		HowTo:   []string{"Yüksek plankta başlayın.", "Elleri normalden yakın koyun.", "Dirsekleri gövdeye yakın tutun.", "Yavaş inip itin."},
		Correct: []string{"Dirsekler yakın.", "Vücut düz.", "Omuzlar kulaklardan uzak.", "Bilekler rahat."},
		Wrong:   []string{"Elleri çok yakın koymak.", "Dirsekleri açmak.", "Beli çökertmek.", "Bilek ya da dirsek ağrısına katlanmak."},
	},
	"P03": {
		Desc:    "Alt noktada kısa duraklamalı kontrollü şınav.",
		HowTo:   []string{"Normal şınav pozisyonunda başlayın.", "Yavaşça inin.", "Altta bir saniye durun.", "Pozisyonu bozmadan itin."},
		Correct: []string{"Duraklama kontrollü.", "Kor sıkı.", "Göğüs aktif.", "Vücut düz kalır."},
		Wrong:   []string{"Altta gevşemek.", "Kalçayı düşürmek.", "Nefesi tutmak.", "Alt noktadan sıçramak."},
	},
	"P04": {
		Desc:    "Kontrolü ve gerilim süresini artıran yavaş şınav çeşidi.",
		HowTo:   []string{"Yüksek plankta başlayın.", "Yaklaşık üç saniye inin.", "Altta bir saniye durun.", "Yaklaşık üç saniye itin."},
		Correct: []string{"Yavaş ve düzenli tempo.", "Vücut düz.", "Dirsekler kontrollü.", "Nefes düzenli."},
		Wrong:   []string{"Acele etmek.", "Altta çökmek.", "Kalçayı düşürmek.", "Korun gerilimini kaybetmek."},
	},
	"P05": {
		Desc:    "Göğüs, omuz ve koru zorlayan asimetrik bir şınav çeşidi.",
		HowTo:   []string{"Şınav pozisyonunda başlayın.", "Bir eli hafif önde, diğerini hafif arkada koyun.", "Kontrolle inin.", "İtin.", "Diğer tarafta el konumunu değiştirin."},
		Correct: []string{"Gövde düzgün.", "Kor sıkı.", "İki omuz da sabit.", "Hareket kontrollü."},
		Wrong:   []string{"Gövdeyi burmak.", "Bir omzu düşürmek.", "Elleri çok geniş açmak.", "Omuz rahatsızsa aralığı zorlamak."},
	},
	"P06": {
		Desc:    "Omuza odaklı bir şınav çeşidi.",
		HowTo:   []string{"Yüksek plankta başlayın.", "Ayakları biraz yaklaştırıp kalçayı kaldırın.", "Baş kollar arasında.", "Başı yere doğru indirin.", "Tekrar itin."},
		Correct: []string{"Kalça yukarıda.", "Dirsekler kontrollü bükülür.", "Omuzlar sabit.", "Boyun nötr."},
		Wrong:   []string{"Normal şınava düşmek.", "Başı çok öne götürmek.", "Dirsekleri çok açmak.", "Omuz ağrısına katlanmak."},
	},
	"P08": {
		Desc:    "Önkol desteği pozisyonundan yapılan triceps ve omuz egzersizi.",
		HowTo:   []string{"Önkollar üzerinde, vücut uzun başlayın.", "Dirsekler omuz altında ya da biraz önde.", "Eller ve önkollarla itin.", "Dirsekleri hafif açın.", "Kontrolle dönün."},
		Correct: []string{"Kor sıkı.", "Dirsekler akıcı.", "Omuzlar aşağıda.", "Hareket aralığı rahat."},
		Wrong:   []string{"Dirsekleri zorlamak.", "Kalçayı düşürmek.", "Omuzları kaldırmak.", "Dirsek ya da omuz ağrısıyla hareket etmek."},
	},

	// ---- Sırt / duruş ----
	"B01": {
		Desc:    "Üst sırt, kürek kemikleri ve omurga kasları için duruş odaklı sırt egzersizi.",
		HowTo:   []string{"Yüzüstü yatın.", "Kollar gövde boyunca.", "Baş parmakları hafif dışa çevirin.", "Göğsü hafif kaldırın.", "Kürek kemiklerini geri ve aşağı çekin.", "Tutun."},
		Correct: []string{"Kalkış küçük.", "Boyun nötr.", "Kürekler geri ve aşağı.", "Kalça hafif aktif."},
		Wrong:   []string{"Çok yükseğe kalkmak.", "Başı geri atmak.", "Beli aşırı çukurlaştırmak.", "Omuzları kulaklara kaldırmak."},
	},
	"B02": {
		Desc:    "Yüzüstü yapılan üst sırt ve kürek kemiği egzersizi.",
		HowTo:   []string{"Yüzüstü yatın.", "Baş aşağı baksın.", "Kolları yanlardan başa doğru geniş yayla götürün.", "Aynı yoldan dönün.", "Hareket kontrollü."},
		Correct: []string{"Boyun gevşek.", "Kürekler akıcı.", "Göğüs kalkışı küçük ya da nötr.", "Kollar acele etmeden."},
		Wrong:   []string{"Başı kaldırmak.", "Beli çok çukurlaştırmak.", "Kolları savurmak.", "Omuzları kulaklara kaldırmak."},
	},
	"B03": {
		Desc:    "Kürek kemikleri ve arka omuz için üst sırt egzersizi.",
		HowTo:   []string{"Yüzüstü yatın.", "Dirsekleri W biçiminde bükün.", "Dirsek ve elleri hafif kaldırın.", "Kürekleri sıkın.", "Kontrolle indirin."},
		Correct: []string{"Küçük hareket aralığı.", "İş kürekler tarafından yapılır.", "Boyun nötr.", "Omuzlar kulaklardan uzak."},
		Wrong:   []string{"Çok yükseğe kaldırmak.", "Başı kaldırmak.", "Momentum kullanmak.", "Beli aşırı çukurlaştırmak."},
	},
	"B04": {
		Desc:    "Alt trapez ve üst sırt için kürek kemiği egzersizi.",
		HowTo:   []string{"Yüzüstü yatın.", "Kolları çapraz, baş üstünde Y biçiminde uzatın.", "Kolları yerden hafif kaldırın.", "Kontrolle indirin."},
		Correct: []string{"Kollar yalnızca biraz kalkar.", "Boyun nötr.", "Omuzlar aşağıda.", "Hareket kontrollü."},
		Wrong:   []string{"Çok yükseğe kaldırmak.", "Omuzları kulaklara kaldırmak.", "Beli çukurlaştırmak.", "Çok hızlı hareket etmek."},
	},
	"B05": {
		Desc:    "Arka omuz ve kürek kontrolü için üst sırt egzersizi.",
		HowTo:   []string{"Yüzüstü yatın.", "Kolları yanlara T biçiminde açın.", "Kolları hafif kaldırın.", "Kürekleri sıkın.", "Yavaşça indirin."},
		Correct: []string{"Kollar omuz hizasında.", "Boyun gevşek.", "Kürekler yumuşak hareket eder.", "Kalkış küçük."},
		Wrong:   []string{"Kolları yukarı savurmak.", "Başı kaldırmak.", "Omuzları kaldırmak.", "Geniş aralığı zorlamak."},
	},
	"B06": {
		Desc:    "Ekipmansız çekiş hareketini taklit eden sırt egzersizi.",
		HowTo:   []string{"Yüzüstü yatın, kollar başın üstünde.", "Göğsü hafif kaldırın.", "Dirsekleri kaburgalara doğru çekin.", "Kürekleri sıkın.", "Kolları tekrar öne uzatın."},
		Correct: []string{"Dirsekler kontrollü iner.", "Kürekler hafif sıkışır.", "Boyun nötr.", "Bel rahat."},
		Wrong:   []string{"Kolları sarsmak.", "Çok yükseğe kalkmak.", "Başı geri atmak.", "Omuzları kulaklara kaldırmak."},
	},
	"B07": {
		Desc:    "Yüzüstü yapılan sırt ve koordinasyon egzersizi.",
		HowTo:   []string{"Yüzüstü yatın.", "Göğsü hafif kaldırın.", "Karşı kol ve bacağı yüzme gibi hareket ettirin.", "Hareket küçük ve düzenli."},
		Correct: []string{"Boyun nötr.", "Hareket kontrollü.", "Bel rahat.", "Nefes sürer."},
		Wrong:   []string{"Çok hızlı hareket etmek.", "Çok yükseğe kalkmak.", "Başı geri atmak.", "Beli zorlamak."},
	},
	"B08": {
		Desc:    "Kalça, arka bacak, omuz ve sırt için arka zincir egzersizi.",
		HowTo:   []string{"Bacaklar uzanmış şekilde yere oturun.", "Elleri kalçanın arkasına koyun.", "El ve topuklarla itin.", "Kalçayı kaldırın.", "Vücudu uzun bir çizgide tutun."},
		Correct: []string{"Göğüs açık.", "Kalça yukarıda.", "Omuzlar sabit.", "Boyun nötr."},
		Wrong:   []string{"Kalçayı düşürmek.", "Omuzları kaldırmak.", "Başı geri atmak.", "Bilekleri zorlamak."},
	},
	"B09": {
		Desc:    "Omurga kasları ve duruş için küçük aralıklı sırt egzersizi.",
		HowTo:   []string{"Yüzüstü yatın.", "Eller gövde yanında ya da göğse yakın.", "Göğsü hafif kaldırın.", "Küçük, kontrollü ritmik hareketler yapın.", "Boyun nötr."},
		Correct: []string{"Küçük kalkış.", "Akıcı ritim.", "Boyun uzun.", "Bel kontrollü."},
		Wrong:   []string{"Çok yükseğe çıkmak.", "Momentum kullanmak.", "Öne bakmak.", "Belde keskin ağrı hissetmek."},
	},
	"B10": {
		Desc:    "Küçük kobra kalkışını W biçimli kol çekişiyle birleştiren daha zorlu sırt egzersizi.",
		HowTo:   []string{"Yüzüstü, kollar önde yatın.", "Göğsü hafif kaldırın.", "Dirsekleri aşağı ve geriye W biçiminde çekin.", "Kürekleri sıkın.", "Tekrar öne uzanıp kontrolle indirin."},
		Correct: []string{"Küçük göğüs kalkışı.", "Akıcı W çekişi.", "Kürekler geri ve aşağı.", "Boyun uzun."},
		Wrong:   []string{"Beli aşırı çukurlaştırmak.", "Boyunla çekmek.", "Çok hızlı hareket etmek.", "Omuzları kaldırmak."},
	},

	// ---- Kollar / biceps ----
	"A01": {
		Desc:    "Bir elin diğerine direnç uyguladığı ekipmansız biceps egzersizi.",
		HowTo:   []string{"Dik durun ya da oturun.", "Bir dirseği curl gibi bükün.", "Diğer elle önkola aşağı bastırın.", "Dirence karşı yavaşça yukarı kaldırın.", "Taraf değiştirin."},
		Correct: []string{"Direnç sabit.", "Hareket yavaş.", "Dirsek gövdeye yakın.", "Omuz aşağıda."},
		Wrong:   []string{"Gerçek direnç olmadan hareket etmek.", "Kolu sarsmak.", "Omzu kaldırmak.", "Nefesi tutmak."},
	},
	"A02": {
		Desc:    "Kendine dirençle statik biceps tutuşu.",
		HowTo:   []string{"Dirseği yaklaşık 90 dereceye bükün.", "Diğer elle aşağı bastırın.", "Çalışan kol hareketsiz direnir.", "Gerilimi tutun.", "Taraf değiştirin."},
		Correct: []string{"Dirsek açısı sabit.", "Gerilim sabit.", "Omuzlar gevşek.", "Nefes sürer."},
		Wrong:   []string{"Kısa sarsıntılarla bastırmak.", "Gövdeyi burmak.", "Omzu kaldırmak.", "Nefesi tutmak."},
	},
	"A03": {
		Desc:    "Ellerin birbirine karşı kullanıldığı izometrik kol ve önkol egzersizi.",
		HowTo:   []string{"Parmakları ya da elleri kenetleyin.", "Dirsekler hafif bükülü.", "Elleri birbirinden uzağa çekin.", "Sabit gerilim tutun.", "Yavaşça gevşeyin."},
		Correct: []string{"Gerilim kontrollü.", "Omuzlar aşağıda.", "Bilekler rahat.", "Nefes sakin."},
		Wrong:   []string{"Sarsarak çekmek.", "Omuzları kaldırmak.", "Bilekleri aşırı bükmek.", "Nefesi tutmak."},
	},
	"A04": {
		Desc:    "Kendine dirençle biceps için yavaş indirme egzersizi.",
		HowTo:   []string{"Bir dirsek bükülü başlayın.", "Diğer elle direnç oluşturun.", "Önkolu yavaşça indirin.", "İndirme boyunca direnci sürdürün.", "Taraf değiştirin."},
		Correct: []string{"İndirme yavaş.", "Direnç sabit.", "Dirsek gövdeye yakın.", "Omuz gevşek."},
		Wrong:   []string{"Kolu hızla düşürmek.", "Direnç kullanmamak.", "Gövdeyi burmak.", "Nefesi tutmak."},
	},

	// ---- Bacaklar / kalça ----
	"L01": {
		Desc:    "Bacaklar, kalça ve duruş için kontrollü bir bacak egzersizi.",
		HowTo:   []string{"Ayaklar kalçadan biraz geniş.", "Kalçayı geriye itin.", "Dizleri rahat, sığ bir derinliğe bükün.", "Yavaşça doğrulun."},
		Correct: []string{"Dizler parmak yönünde.", "Topuklar yerde.", "Sırt uzun.", "Derinlik rahat."},
		Wrong:   []string{"Çok derine inmek.", "Dizleri içe düşürmek.", "Topukları kaldırmak.", "Hızla çökmek."},
	},
	"L02": {
		Desc:    "Güvenli aralıkta kısa duraklamalı kontrollü çömelme çeşidi.",
		HowTo:   []string{"Ayaklar sabit durun.", "Rahat, sığ bir çömelmeye inin.", "Kısa süre durun.", "Kontrolle doğrulun."},
		Correct: []string{"Duraklama sabit.", "Dizler parmak yönünde.", "Göğüs açık.", "Topuklar yerde."},
		Wrong:   []string{"Çok derine inmek.", "Altta zıplamak.", "Dizleri içe düşürmek.", "Nefesi tutmak."},
	},
	"L03": {
		Desc:    "Kalça, arka bacak ve bel kontrolü için bir hareket çalışması.",
		HowTo:   []string{"Dik durun, ayaklar kalça genişliğinde.", "Dizleri hafif yumuşatın.", "Kalçayı geriye itin.", "Gövdeyi düz sırtla eğin.", "Kalçayı öne iterek doğrulun."},
		Correct: []string{"Sırt düz.", "Hareket kalçadan başlar.", "Dizler yumuşak, çok bükülmemiş.", "Boyun omurga çizgisini izler."},
		Wrong:   []string{"Sırtı yuvarlamak.", "Çömelmeye çevirmek.", "Çok aşağı uzanmak.", "Başı çok kaldırmak."},
	},
	"L04": {
		Desc:    "Tek bacakla yapılan daha zorlu bir kalça egzersizi.",
		HowTo:   []string{"Sırtüstü yatın.", "Bir dizi bükün, ayak yerde.", "Diğer bacağı uzatın ya da kaldırın.", "Çalışan topuktan itin.", "Kalçayı kaldırıp yavaşça indirin."},
		Correct: []string{"Kalkış çalışan kalçadan.", "Kalça düz.", "Hareket kontrollü.", "Bel nötr."},
		Wrong:   []string{"Pelvisi burmak.", "Parmak ucundan itmek.", "Beli çukurlaştırmak.", "Hızla düşmek."},
	},
	"L05": {
		Desc:    "Kalça ve pelvis stabilitesi egzersizi.",
		HowTo:   []string{"Köprü pozisyonunda başlayın.", "Kalçayı yukarıda tutun.", "Bir ayağı yerden hafif kaldırın.", "İndirip taraf değiştirin.", "Pelvis sabit."},
		Correct: []string{"Kalça düz.", "Kalça kasları aktif.", "Hareket yavaş.", "Bel rahat."},
		Wrong:   []string{"Kalçayı düşürmek.", "Yana sallanmak.", "Çok hızlı hareket etmek.", "Beli çukurlaştırmak."},
	},
	"L06": {
		Desc:    "Kalça ve arka zinciri güçlendiren bir kalça egzersizi.",
		HowTo:   []string{"Sırtüstü yatın.", "Dizleri bükün, ayaklar yerde.", "Kalçayı kaldırın.", "Üstte durun.", "Yavaşça indirin."},
		Correct: []string{"Topuklardan itin.", "Üstte kalça sıkışır.", "Kaburgalar aşağıda.", "Bel aşırı açılmaz."},
		Wrong:   []string{"Üstte beli çukurlaştırmak.", "Topukları kaldırmak.", "Kalçayı çok hızlı düşürmek.", "Ayakları çok yakın koymak."},
	},
	"L07": {
		Desc:    "Baldır, ayak bileği ve ayak kontrolü için bir alt bacak egzersizi.",
		HowTo:   []string{"Dik durun, ayaklar kalça genişliğinde.", "Parmak uçlarına yükselin.", "Üstte kısa durun.", "Topukları yavaşça indirin."},
		Correct: []string{"Vücut dik.", "Hareket akıcı.", "Ayak bilekleri hizalı.", "İndirme kontrollü."},
		Wrong:   []string{"Topukları hızla düşürmek.", "Öne sallanmak.", "Ayak bileklerini dışa devirmek.", "Momentum kullanmak."},
	},
	"L09": {
		Desc:    "Yan kalça, yan karın ve kalça stabilitesi için yanal zincir egzersizi.",
		HowTo:   []string{"Yan plankta başlayın.", "Vücut düz çizgide.", "Üstteki bacağı hafif kaldırın.", "Kontrolle indirin.", "Tekrarlayıp taraf değiştirin."},
		Correct: []string{"Kalça yukarıda.", "Üst bacak yavaş hareket eder.", "Gövde sabit.", "Omuz güçlü."},
		Wrong:   []string{"Kalçayı düşürmek.", "Bacağı savurmak.", "Geriye yuvarlanmak.", "Omza çökmek."},
	},

	// ---- Soğuma ----
	"CD01": {
		Desc:    "Göğüs ve omuz önü için nazik bir esneme.",
		HowTo:   []string{"Sırtüstü yatın.", "Kolları yanlara açın.", "Omuzları gevşetin.", "Göğsün nazikçe açılmasına izin verin.", "Yavaş nefes alın."},
		Correct: []string{"Göğüs yumuşak açılır.", "Omuzlar gevşek.", "Bel rahat.", "Omuzda ağrı yok."},
		Wrong:   []string{"Kolları zorla bastırmak.", "Beli çukurlaştırmak.", "Boynu kasmak.", "Ağrıya rağmen esnetmek."},
	},
	"CD02": {
		Desc:    "Uyluk arkası için nazik bir esneme.",
		HowTo:   []string{"Sırtüstü yatın.", "Bir bacağı kaldırın.", "Uyluk ya da baldırın arkasından tutun.", "Bacağı nazikçe kendinize çekin.", "Taraf değiştirin."},
		Correct: []string{"Esneme hafif.", "Diz hafif bükülü kalabilir.", "Bel sakin.", "Nefes yavaş."},
		Wrong:   []string{"Çok sert çekmek.", "Bacağı zorla düzleştirmek.", "Kalçayı kaldırmak.", "Diz arkasında keskin ağrı hissetmek."},
	},
	"CD03": {
		Desc:    "Omurga ve gövde için nazik bir dönüş esnemesi.",
		HowTo:   []string{"Sırtüstü yatın.", "Dizleri bükün.", "Dizleri nazikçe bir yana bırakın.", "Omuzları yerde tutun.", "Ortaya dönüp taraf değiştirin."},
		Correct: []string{"Hareket yavaş.", "Omuzlar yerde.", "Nefes gevşek.", "Esneme rahat."},
		Wrong:   []string{"Dizleri zorla yere bastırmak.", "Hızlı dönmek.", "Omzu kaldırmak.", "Ağrıya rağmen esnetmek."},
	},
	"CD04": {
		Desc:    "Nazik bir ön gövde esnemesi ve hafif sırt uzaması.",
		HowTo:   []string{"Yüzüstü yatın.", "Önkolları yere koyun.", "Göğsü nazikçe kaldırın.", "Boynu uzun tutun.", "Yavaş nefes alın."},
		Correct: []string{"Kalkış hafif.", "Omuzlar kulaklardan uzak.", "Bel rahat.", "Boyun nötr."},
		Wrong:   []string{"Beli aşırı çukurlaştırmak.", "Başı geri atmak.", "Omuzları kaldırmak.", "Pozisyonu zorlamak."},
	},
	"CD05": {
		Desc:    "Seansı bitiren sakin bir nefes egzersizi.",
		HowTo:   []string{"Sırtüstü yatın.", "Omuzları ve çeneyi gevşetin.", "Rahatsa bir eli karnınıza koyun.", "Nazikçe nefes alın.", "Yavaşça verin."},
		Correct: []string{"Nefes sakin.", "Omuzlar gevşek.", "Yüz yumuşak.", "Vücut sakinleşir."},
		Wrong:   []string{"Çok güçlü nefes almak.", "Nefesi tutmak.", "Boynu kasmak.", "Beli çukurlaştırmak."},
	},

	// ---- Vlad seti eklemeleri (ısınma / kardiyo / pliometri / hamleler) ----
	"W07": {
		Desc:    "Omurgayı uzatmak için nazik bir tam vücut uzanmasıyla derin nefes alan sakin bir açılış çalışması.",
		HowTo:   []string{"Dik durun, ayaklar kalça genişliğinde.", "Nefes alın ve iki kolu başın üstüne uzatın.", "Omurga boyunca nazikçe uzanın.", "Nefes verin ve kolları indirip omuzları bırakın."},
		Correct: []string{"Nefes yavaş ve dolu.", "Nefes verirken omuzlar gevşer.", "Kaburgalar aşağıda, aşırı çukurlaşma yok.", "Hareket acelesiz."},
		Wrong:   []string{"Nefesi tutmak.", "Beli sertçe çukurlaştırmak.", "Omuzları kulaklara kaldırmak.", "Uzanmayı acele etmek."},
	},
	"W08": {
		Desc:    "Omurgayı ve gövdeyi ısıtan nazik bir rotasyon çalışması.",
		HowTo:   []string{"Ayaklar kalça genişliğinde durun.", "Kolların serbestçe sallanmasına izin verin.", "Üst gövdeyi bir yana çevirin.", "Akıcı şekilde diğer yana geçin."},
		Correct: []string{"Hareket gövdeden gelir.", "Kalça çoğunlukla öne bakar.", "Ayaklar sabit.", "Tempo akıcı ve düzenli."},
		Wrong:   []string{"Dizleri sertçe burmak.", "Kolları savurmak.", "Çok hızlı dönmek.", "Nefesi tutmak."},
	},
	"C13": {
		Desc:    "Karın ve yan karın için dinamik bir kor egzersizi.",
		HowTo:   []string{"Sırtüstü yatın, eller hafifçe başın arkasında.", "Omuzları kaldırın ve dizleri yukarı getirin.", "Bir dirseği karşı dize doğru götürün.", "Düzenli bir pedal çevirme hareketiyle taraf değiştirin."},
		Correct: []string{"Bel yerde kalır.", "Hareket kontrollü, savrulmaz.", "Dirsekler geniş.", "Nefes düzenli."},
		Wrong:   []string{"Boyundan çekmek.", "Beli yerden kaldırmak.", "Ritmi acele etmek.", "Nefesi tutmak."},
	},
	"B11": {
		Desc:    "Karşı kol ve bacağı almaşık, yüzme benzeri bir hareketle kaldıran sırt egzersizi.",
		HowTo:   []string{"Yüzüstü yatın, kollar başın üstünde uzanmış.", "Göğsü ve bacakları hafifçe kaldırın.", "Bir kolu ve karşı bacağı kaldırın.", "İndirip taraf değiştirin, akıcı şekilde almaşık yapın."},
		Correct: []string{"Kalkış küçük kalır.", "Boyun omurga çizgisini izler.", "Kalça kasları aktif.", "Hareket düzenli."},
		Wrong:   []string{"Beli zorlamak.", "Başı kaldırmak.", "Sarsıntılı hareket etmek.", "Nefesi tutmak."},
	},
	"L10": {
		Desc:    "Uyluk önü, kalça kasları ve kalça için temel bir bacak egzersizi.",
		HowTo:   []string{"Ayaklar omuz genişliğinde durun.", "Kalçayı geriye itin ve dizleri bükün.", "Uyluklar paralele yakın olana dek inin.", "Topuklardan iterek doğrulun."},
		Correct: []string{"Dizler parmak yönünde.", "Topuklar yerde.", "Göğüs açık.", "Sırt nötr."},
		Wrong:   []string{"Dizleri içe düşürmek.", "Topukları kaldırmak.", "Sırtı yuvarlamak.", "Kontrolsüz çökmek."},
	},
	"L11": {
		Desc:    "Bacaklar ve kalça için dize nazik, tek bacaklı bir egzersiz.",
		HowTo:   []string{"Dik durun, ayaklar kalça genişliğinde.", "Bir ayağı geriye adımlayın.", "İki diz de 90 dereceye yakın olana dek inin.", "Ön topuktan iterek dönün."},
		Correct: []string{"Ön diz topuğun üstünde kalır.", "Gövde dik.", "Adım kontrollü.", "Arka diz yumuşakça iner."},
		Wrong:   []string{"Ön dizi içe düşürmek.", "Gövdeyi çok öne eğmek.", "Arka dizi yere çarpmak.", "Adımları acele etmek."},
	},
	"L12": {
		Desc:    "Kontrol ve bacak gücü geliştirmek için yavaş yapılan bir çömelme.",
		HowTo:   []string{"Ayaklar omuz genişliğinde durun.", "Yaklaşık üç saniyede inin.", "Altta kısa süre durun.", "Kontrolle doğrulun."},
		Correct: []string{"Tempo yavaş kalır.", "Dizler parmak yönünde.", "Topuklar yerde.", "Sırt nötr."},
		Wrong:   []string{"Hızla çökmek.", "Alt noktadan sıçramak.", "Dizleri içe düşürmek.", "Nefesi tutmak."},
	},
	"J01": {
		Desc:    "Baldırları ve ayak bileklerini hazırlamak için ayak uçlarından küçük, yaylı sıçramalar.",
		HowTo:   []string{"Dik durun, ayaklar bitişik.", "Ayak uçlarından hafifçe zıplayın.", "Dizleri hafif yumuşak tutun.", "Yumuşak ve sessizce inin."},
		Correct: []string{"Sıçramalar alçak kalır.", "İnişler sessiz.", "Dizler yumuşak.", "Vücut dik."},
		Wrong:   []string{"Sert ve gürültülü inmek.", "Dizleri kilitlemek.", "Çok yükseğe zıplamak.", "Ayak bileklerini içe çökertmek."},
	},
	"J02": {
		Desc:    "Dizleri kaldıran düşük etkili bir yürüyüş, nazik bir kardiyo ısınması.",
		HowTo:   []string{"Dik durun.", "Yerinde yürüyün, bir dizi kalça hizasına kaldırın.", "Karşı kolu doğal şekilde sallayın.", "Düzenli, rahat bir tempo tutun."},
		Correct: []string{"Gövde dik.", "Dizler rahat bir yüksekliğe kalkar.", "Adımlar hafif.", "Nefes rahat."},
		Wrong:   []string{"Geriye yaslanmak.", "Ayakları yere vurmak.", "Nefesi tutmak.", "Kontrolü kaybedecek kadar acele etmek."},
	},
	"J03": {
		Desc:    "Tüm vücudu ısıtan klasik bir tam vücut kardiyo hareketi.",
		HowTo:   []string{"Ayaklar bitişik, kollar yanlarda durun.", "Kollar başın üstüne kalkarken ayakları yana açın.", "Başlangıca geri zıplayın.", "Düzenli bir ritim tutun."},
		Correct: []string{"İnişler yumuşak kalır.", "Dizler parmak yönünde.", "Hareket ritmik.", "Nefes düzenli."},
		Wrong:   []string{"Sert inmek.", "Dizleri içe düşürmek.", "Omuzları kamburlaştırmak.", "Nefesi tutmak."},
	},
	"J04": {
		Desc:    "Plank pozisyonundan dizleri göğse doğru süren dinamik bir kor ve kardiyo hareketi.",
		HowTo:   []string{"Yüksek plankta başlayın, eller omuzların altında.", "Bir dizi göğse doğru sürün.", "Bacakları hızlıca değiştirin.", "Kalçayı alçak ve sabit tutun."},
		Correct: []string{"Kalça hizada kalır.", "Omuzlar ellerin üstünde.", "Kor sıkı.", "Tempo kontrollü."},
		Wrong:   []string{"Kalçayı yukarı zıplatmak.", "Beli çökertmek.", "Omuzları çok öne kaydırmak.", "Nefesi tutmak."},
	},
	"J05": {
		Desc:    "Patlayıcı bir alt vücut hareketi: çömelip sonra güçlüce yukarı zıplamak.",
		HowTo:   []string{"Ayaklar omuz genişliğinde durun.", "Bir çömelmeye inin.", "Patlayıcı şekilde yukarı zıplayın.", "Yumuşak inip bir sonraki çömelmeye geçin."},
		Correct: []string{"İnişler yumuşak ve sessiz.", "Dizler parmak yönünde.", "Göğüs açık.", "Kalça inişi emer."},
		Wrong:   []string{"Sert ve gürültülü inmek.", "Dizleri içe düşürmek.", "Sırtı yuvarlamak.", "Düz bacaklarla inmek."},
	},
	"J06": {
		Desc:    "Güç, denge ve koordinasyon için yandan yana yanal bir sıçrama.",
		HowTo:   []string{"Tek bacak üzerinde durun.", "Yana doğru diğer bacağa zıplayın.", "Arka bacağın arkadan sallanmasına izin verin.", "Yumuşak inip ters yöne yapın."},
		Correct: []string{"İnişler yumuşak kalır.", "Diz ayağın üstünde kalır.", "Kalça dengeli.", "Hareket akıcı."},
		Wrong:   []string{"Dizi içe düşürmek.", "Sert inmek.", "Dengeyi kaybetmek.", "Kontrolsüz acele etmek."},
	},
	"J07": {
		Desc:    "Dizleri yükseğe süren, yerinde koşma kardiyo hareketi.",
		HowTo:   []string{"Dik durun.", "Yerinde koşun, dizleri kalça hizasına kaldırın.", "Ayak uçlarında kalın.", "Kolları ritimle çalıştırın."},
		Correct: []string{"Gövde dik.", "Dizler yükseğe kalkar.", "İnişler hafif.", "Tempo düzenli."},
		Wrong:   []string{"Geriye yaslanmak.", "Ayakları yere vurmak.", "Kamburlaşmak.", "Nefesi tutmak."},
	},
	"J08": {
		Desc:    "Çömelme, plank ve sıçramayı birleştiren tam vücut kondisyon hareketi.",
		HowTo:   []string{"Ayaktayken çömelip elleri yere koyun.", "Ayakları zıplayarak ya da adımlayarak planka atın.", "Ayakları kalçanın altına geri getirin.", "Doğrulup uzanarak zıplayın."},
		Correct: []string{"Plankta omurga nötr.", "İnişler yumuşak kalır.", "Tempo düzenli.", "Kor aktif."},
		Wrong:   []string{"Plankta kalçayı çökertmek.", "Sert inmek.", "Formu bozacak kadar acele etmek.", "Nefesi tutmak."},
	},
	"J09": {
		Desc:    "Güç ve koordinasyon için havada bacak değiştiren patlayıcı bir hamle.",
		HowTo:   []string{"Hamle pozisyonunda başlayın.", "Yukarı zıplayıp havada bacak değiştirin.", "Diğer tarafta yumuşak bir hamleye inin.", "Almaşık şekilde sürdürün."},
		Correct: []string{"İnişler yumuşak kalır.", "Ön diz ayağın üstünde kalır.", "Gövde dik.", "Kalça her inişi emer."},
		Wrong:   []string{"Sert inmek.", "Ön dizi içe düşürmek.", "Çok öne eğilmek.", "Değişimin kontrolünü kaybetmek."},
	},
	"CD07": {
		Desc:    "Sakinleşmek için sırt, kalça ve omuzlara dinlendirici bir esneme.",
		HowTo:   []string{"Diz çöküp kalçayı topuklara doğru geri oturtun.", "Kolları yerde öne uzatın.", "Alnın yere dayanmasına izin verin.", "Yavaşça nefes alıp gevşeyin."},
		Correct: []string{"Kalça topuklara doğru yerleşir.", "Sırt nazikçe uzar.", "Omuzlar gevşer.", "Nefes yavaş kalır."},
		Wrong:   []string{"Kalçayı zorla aşağı bastırmak.", "Omuzları kasmak.", "Nefesi tutmak.", "Dizleri zorlamak."},
	},
}
