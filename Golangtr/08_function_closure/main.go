package main

import "fmt"

/*
Closure(anonim) fonksiyonlar
closure fonksiyonlar ile değişkenlerimizi fonksiyon olarak tanımlayabilir ve kullanabiliriz.
*/
func main() {
	toplam := func(x, y, z int) (toplam int) {
		toplam = x + y + z
		return
	}

	fmt.Println(toplam(3, 4, 5))

	// variadic olarak aynı işlemi yapalım.

	toplam2 := func(sayilar ...int) (toplam int) {
		for _, sayi := range sayilar {
			toplam += sayi
		}
		return
	}

	fmt.Println(toplam2(1, 2, 3, 4, 5, 6, 7, 8, 9))
}
