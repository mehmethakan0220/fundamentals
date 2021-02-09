package main

import "fmt"

/*
Anonim fonksiyonların en büyük üzelliği isimsiz olmalarıdır. Yazıldıkları yerde direkt olarak çalışırlar.
Çalışırken diğer fonksiyonlar gibi parametre  verilmediği  için fonksiyonun sonuna arguman eklenerek çalıştırılırlar
*/

func main() {

	//anonim fonksiyon
	func(isim string) {
		fmt.Println(isim)
	}("harun") // burada argümanı direk veriyoruz fonksiyon olduğu yerde çalışıyor

}
