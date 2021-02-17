package main

import "fmt"

func main() {

	var harita map[string]int // define

	harita = make(map[string]int) // initialize

	harita["sayi1"] = 234
	harita["sayi2"] = 432
	harita["sayi3"] = 23

	fmt.Println(harita)

	//şimdi 23 ü haritadan silelim :D

	delete(harita, "sayi3") // silme işlemi için delete kullanılır.
	fmt.Println(harita)
}
