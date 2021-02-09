package main

/*
Bu bölümde primitive tipler üzerinde yeni tipler oluşturup bu tiplere fonsiyon yazacağız
*/
import (
	"fmt"
)

type katar string // string türünde katar adında yeni bir tip oluşturduk

func (k katar) uzunluk() int { //katar tipine ait uzunluk diye bir fonksiyon oluşturduk.
	return len(k)
}

type dordislem int

func (d dordislem) topla(sayilar ...int) (sonuc int) {
	for _, sayi := range sayilar {
		sonuc += sayi
	}
	return
}

func main() {
	isim := katar("MehmetHakan") // katar tipinden örnek oluşturduk.
	fmt.Println("Ad:", isim, " ", "isim uzunluğu:", isim.uzunluk())

	islem := dordislem(0) //burada hata vermesin diye 0 verdik önemi yok

	fmt.Println(islem.topla(1, 2, 3, 4, 5, 6, 7, 8, 9))

}
