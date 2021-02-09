package main

import "fmt"

func main() {

	//not

	// birşeye anonim diyor isek oluşturulduğu yerde initialize edilmesi gereklidir.

	//tıpkı anonim fonksiyonlarda olduğu gibi

	kisi := struct { // define işlemi
		ad, soyad string
	}{"kemal", "tanca"} //initialize işlemi

	fmt.Println(kisi)

	//anonim bir fonksiyon

	myfank := func(x, y int) int { //define islemi
		return x + y
	}(5, 6) // initialize işlemi

	fmt.Println(myfank)

	//kopmlex bir anonim map oluşturalım. struct yapısı kullanarak.

	mixed := map[int]struct {
		marka, model string
		yil          int
	}{
		1: {"bmw", "x6", 2020},
		2: {"tofaş", "kartal", 95},
	}

	fmt.Println(mixed)
	// adım adım ne yaptık

	// [int] yazılı olan kısım map ile alakalı orayı görmezden gelirsek
	// struct {} ifadesinde yapımızın alacağı özellikleri yazıyoruz.
	// yani map[int] kısmının struct yapısı ile işi yok
	// [int] map ile alakalı olduğu için initialize ederken ora için de değer göndermemiz gerekecektir.
}
