package main

import (
	"fmt"
)

type toplam struct {
	sayi1, sayi2 int
}

func (t *toplam) hesapla() {
	fmt.Println(t.sayi1 + t.sayi2)
}

type carpim struct {
	sayi1, sayi2 int
}

func (c *carpim) hesapla() {
	fmt.Println(c.sayi1 * c.sayi2)
}

//şimdi interface imizi oluşturalım.
type arayüz interface {
	hesapla()
}

func sonuc(a arayüz) {
	a.hesapla()
}

//sıradan bir fonksiyon oluşturalım.

func main() {
	islem1 := toplam{3, 5}
	arayüz.hesapla(&islem1) // bu normal interface kullanımı 1 sonraki adıma geçelim
	//bu kullanımıda sonuc fonksiyonunu kulllanmadık.

	islem2 := carpim{9, 7}
	fmt.Println(islem2)
	//şimdi arayüzden bir nesne oluşturralım.
	var nesne arayüz // nill nesne olduğu için bunun initialize edilmesi gerekir.
	if nesne == nil {
		fmt.Println("nesne nil mi?:", nesne)
	}

	var nesne1 arayüz = arayüz(&islem2)

	//şimdi nesne1 i initialize ettik

	nesne1.hesapla() //nesne 1 için hesaplama yapacaktır.

	// bunların hiçbirini istemiyoruz.

	nesne3 := &islem2
	sonuc(nesne3)
	///daha basit bu şekilde kullanılabillir.

}
