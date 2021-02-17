package main

import "fmt"

type insan struct {
	name string
	age  int
}

func main() {

	var kisi map[string]insan = map[string]insan{"personel1": {name: "ali", age: 25}} // tanımladık ve atadık.
	fmt.Println(kisi)
	//bir eleman ekleyelim.

	kisi["personel2"] = insan{name: "suat", age: 24}
	fmt.Println(kisi)

	//make ile boş bir map oluşturmayı deneyelim. ve nill testine koyalım.
	var kisi1 map[string]insan
	if kisi1 == nil {
		fmt.Println(kisi1, "   :nil")
	}
	kisi1 = make(map[string]insan) // zero değer alan bir map oluşturma //boş map //initialize
	if kisi1 == nil {
		fmt.Println(kisi1, "   :nil")
	} else {
		fmt.Println(kisi1, "   :zero values")
	}

	//sadece define edersek nill olur.
	//make ile initialize edersek te zero valuesleri alır.

	//şimdi kisi1 e eleman ekleyelim.

	kisi1["isci1"] = insan{name: "necdet", age: 45}

	fmt.Println(kisi1)
	// bi tane daha

	kisi1["isci2"] = insan{name: "derya", age: 30}
	fmt.Println("kisi1 in son hali:", kisi1)
	// çoklu eleman ekleme.

	kisi1 = map[string]insan{
		"coklu1": {name: "ahmet", age: 25},
		"coklu2": {name: "ferhat", age: 45},
		"coklu3": {name: "musa", age: 20},
	}

	fmt.Println("coklu eleman ekledikten sonra", kisi1)

}
