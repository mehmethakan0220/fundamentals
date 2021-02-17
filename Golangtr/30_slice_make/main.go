package main

import "fmt"

func main() {

	//:= ile

	a := make([]int, 5) // çıktı = [0 0 0 0 0]
	fmt.Println(a)

	// var ve tanımlayıcı kullanarak

	var b []int = make([]int, 4) //  çıktı = [0 0 0 0]
	fmt.Println(b)

	//var ile tanımlayıcı kullanmadan oluşturma

	var c = make([]int, 3) //  çıktı = [0 0 0 0]
	fmt.Println(c)

	fmt.Println("uzunluk:", len(a), " ", "kapasite:", cap(a))

	// şimdiye kadar make hem uzunluğu hem de kapasiteyi verdiğimiz sayıya göre otomatik oluşturuyor.
	//şimdi make ile uzunluğu ve kapasitesi olan slicelar oluşturalım.

	//slice:= make([]int,lenght,capacity)

	d := make([]int, 2, 4) // uzunluğu 2 kapasitesi 4 olan dizi oluşturmaya yarar
	fmt.Println(d)
	fmt.Println("uzunluk:", len(d), " ", "kapasite:", cap(d))

}
