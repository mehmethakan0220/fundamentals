package main

import "fmt"

func main() {
	a := [8]int{7, 6, 5, 4, 3, 2, 1, 0}
	//bir slice ile a dizisini baştan sona dilimleyelim.

	dilim := a[:]

	fmt.Printf("type:	%T,   values:	%v\n", dilim, dilim)

	//dilimlerin uzunluk(len) ve kapasite(cap) değeri vardır.

	fmt.Println("dilim uzunluğu:", len(dilim), " ", "dilim kapasitesi:", cap(dilim))

	//şimdi yeni bir slice oluşturalım.

	dilim1 := a[0:4] //dizinin ilk 4 elemanını dilimleyen slice
	fmt.Println("dilim1 uzunluğu:", len(dilim1), " ", "dilim1 kapasitesi:", cap(dilim1))

	dilim2 := a[4:]
	fmt.Println("dilim2 uzunluğu:", len(dilim2), " ", "dilim2 kapasitesi:", cap(dilim2))

	//bu çıktılardan anlaşılan dilimler diziler üzerinde pointer gibi davranır ve pointerlarda geriye sayma olmadığı için slice ın başlangıç değeri verdiğiniz indexten sonrası slice ın kapasitesini belirler.

}
