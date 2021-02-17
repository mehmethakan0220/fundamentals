package main

import "fmt"

type toplam struct {
	sayi1, sayi2 int
} // bu yapıya ait fonksiyonumuzu da yazalım.
func (t *toplam) hesapla() {
	fmt.Println(t.sayi1 + t.sayi2)
}

type carpim struct {
	sayi1, sayi2 int
} //bu yapıya ait fonksiyonumuzu da yazalım.
func (c *carpim) hesapla() {
	fmt.Println(c.sayi1 * c.sayi2)
}

type hesap interface {
	hesapla()
}

func main() {
	islem1 := toplam{3, 5}
	islem2 := carpim{4, 5}

	hesap.hesapla(&islem1) // interface in hesapla fonksiyonuna islem1 in adresini yolladık
	hesap.hesapla(&islem2)

	//ama biz interfaceteki hesapla fonksiyonunu hesap.hesapla() olarak değilde direk hesapla() olarak çağırmak istiyoruz.
	//bu durumda interface türündeki hesaptan bir nesne türetelim.

}
