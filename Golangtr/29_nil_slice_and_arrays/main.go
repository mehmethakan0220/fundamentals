package main

import "fmt"

func main() {

	var dizi [5]int //tanımladık initialize etmedik

	var slice []int //tanımladık initialize etmedik

	fmt.Println(dizi)  //initialize etmememize rahmet içine 0 ları doldurdu.
	fmt.Println(slice) //boş bir [] gösterdi

	if slice == nil {
		fmt.Println("slice nill")
	}

	// sliceları oluşturup değer atamadığımızda dizilerden farklı olarak nil değerler alırlar.
}
