package main

import "fmt"

type toplam struct {
	sayi1 int
	sayi2 int
} //bu yapıya ait toplama fonksiyonumuzu yazalım.
func (t *toplam) hesapla() {
	fmt.Println(t.sayi1 + t.sayi2)
}

func main() {
	islem1 := toplam{3, 5} // topla tipinde islem1 nesnesi oluşturup değerlerini atadık.
	islem1.hesapla()

	//boş bi tane oluşturmak istersek ne yapacağız.

	islem2 := toplam{} //bu!
	fmt.Println(islem2)
}
