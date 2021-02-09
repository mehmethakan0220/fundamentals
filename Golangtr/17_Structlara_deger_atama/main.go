package main

import "fmt"

/*
Bu bölümde structlara temel olarak 3 farklı şekilde değer ataması göreceğiz
Ayriyeten struct tabanlı yeni türler için de inceleyeceğiz.
*/

var insan struct { // struct tabanlı insan yapısı
	name    string
	surname string
	can     int
}

type hayvan struct { //struct tabanlı bir hayvan tipi
	name string
	can  int
}

func main() {

	// şimdi insan yapısı tanımlandı fakat varsayılan olarak zero değerler aldı yani `boşstring`,`boşstring`,0 şeklinde

	//şimdi insana varsayılan değerleri ekleyeceğiz.

	insan.name = "isim"
	insan.surname = "soyisim"
	insan.can = 1
	//şimdi insanı print ettirelim.

	fmt.Println(insan) // {isim soyisim 1}

	//şimdi bu insan örneğinden yeni insanlar oluşturalım. zero values mi alıyor yoksa bizim verdiğimiz değerleri mi alacak.

	yeniinsan := insan

	fmt.Println(yeniinsan) // {isim soyisim 1} şeklinde bi çıktı verdi yani biz o yapıyı yukarıda doldurduk.
	//yeni ve boş örnekler üretmek istiyorsak önce insan yapısından bir örnek alıp o örneği özelleştirmeliyiz.

	// bundan sonra hayvan ile devam edelim.
	// 1. atama şekli
	kedi := hayvan{
		name: "Mırıl",
		can:  9,
	}
	fmt.Println(kedi)

	// 2. atama şekli
	kopek := hayvan{"karabaş", 2}
	fmt.Println(kopek)

	// 3. atama şekli
	gopher := hayvan{}
	gopher.name = "köstek"
	gopher.can = 1

	// 4. atama şekli 2 ile aynı

	yilan := hayvan{
		"Piton",
		5,
	}
	fmt.Println(yilan)

	// 5. atama şekli 1 ile aynı
	sincap := hayvan{name: "sino", can: 8}
	fmt.Println(sincap)

	// hepsinde önemlisi :D

	var leylek hayvan   // be şekilde görebilmek.
	fmt.Println(leylek) // zero değerleri gösterir.
	leylek.name = "leylekey"
	leylek.can = 0

	//bide define ettiğimiz yerde çalıştırabilsek tamam!

	var yengec hayvan = hayvan{name: "ısırgaç", can: 7} //that's funny
	fmt.Println(yengec)

	//pislik yapacam

	var ayi hayvan = hayvan{
		"yogi",
		12,
	}
	fmt.Println(ayi)
	//birazda yukarıdakini zorlayak insan yapısından element uydurmaya çalışlım.

	//var necati insan = insan{name: "neco", "beko", 8}
	//fmt.Println(necati)

	// yukarıdaki kodların hata vermesinin sebenini bulalım.

	var sayi []int = []int{2, 3, 5, 7, 11, 13}
	fmt.Println(sayi)
	// bu olduysa 89 ile 92. satırların da çalışması lazım.

	/* 	var necati insan = insan{insan.name: "neco", insan.surname: "beko", insan.can: 8}
	   	fmt.Println(necati)
	   	insan bir tip depil diyor
	*/

	/* var kamil insan = insan{{"kemal", "sayar", 12}}
	olmadı :D
	*/

}
