package main

import "fmt"

func bolme(bolunen, bolen int) (bolum, kalan int) {
	bolum = bolunen / bolen
	kalan = bolunen % bolen
	return
}

func main() {
	bolum, kalan := bolme(13, 3)
	fmt.Println("bolum:", bolum, " kalan:", kalan)
}

/*

bolme fonksiyonunda geriye bolum ve kalan şeklinde iki tane veri çıkartılacağını biliyoruz

Bu durumda return keywordunu olduğu gibi kullanabiliriz.

Yani fonksiyon işin bitirdikten sonra geri döndereceği değerleri biliyor ekstradan belirtmeye gerek yok.

*/
