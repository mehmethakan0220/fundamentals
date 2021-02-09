package main

import "fmt"

/*
Golang te genel olarak 3 çeşit fonksiyon bulunmaktadır.
*/

// 1 Variadic fonksiyonlar
/*Variadic fonksiyon tipi ile  fonksiyonumuza  kaç tane değer girişi olduğunu belirtmeden istediğiniz kadar
değer girebilirsiniz*/

func toplama(sayilar ...int) (toplam int) {

	for _, sayi := range sayilar {
		toplam += sayi
	}
	return
}

func main() {
	sonuc := toplama(1, 2, 3, 4, 5, 6, 7, 8, 9, 10) //variadic fonksiyona argüman gönderme
	fmt.Println(sonuc)

}
