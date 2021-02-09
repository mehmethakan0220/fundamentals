package main

import "fmt"

/*
Golang kodlarımızda bazen 2 adet değer döndüren fonksiyonlar kullanırız. Bu değerlerden hangisini kullanmak istemiyorsak, değişken adı yerine _ (alt tire) kullanırız.
*/

func dortis(x, y int) (toplam int, fark int, carpim int) {
	toplam = x + y
	fark = x - y
	carpim = x * y
	return
}

func main() {
	toplam, fark, carpim := dortis(3, 4)

	fmt.Println("toplam", toplam, "fark", fark, "carpim", carpim)

	//şimdi ayni fonksiyonu 5 ve 6 değerleri için çağıralım ve hangi sonucları almak istediğimizi deneyelim.

	toplam1, _, _ := dortis(5, 6)
	fmt.Println("toplam1", toplam1)
	/*
		fonksiyondan dönen ifadeler arasından kullanmak istemediğimizi _ ile görmezden gelebiliriz.
	*/
}
