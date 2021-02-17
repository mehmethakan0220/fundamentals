package main

import "fmt"

func main() {

	//sliceları dizileri dilimleme işlemi olarak kullanabiliriz.

	var dizi [5]int = [5]int{0, 1, 2, 3, 4} //dizimizi eşitliğin sol tarafında tanımlayıcısı ile atadık.
	fmt.Println(dizi)
	var dizi1 = [5]int{5, 6, 7, 8, 9} // dizimizi eşitliğin sol tarafında tanımlayıcı kullanmadık.
	fmt.Println(dizi1)
	//şimdi burada slice ile dizimizi dilimleyerek alalım.
	//slice tanımlamak için [] sadece köşeli parantezleri kullanırız

	//var slice_ismi []slice_türü = dilimlenece_dizi_adi[start:end]
	var slice []int = dizi[1:3] // [1,3)  yani 1.indexten başlar 3 indexe kadar gider 3 dahil değil
	fmt.Println("dilimimiz:", slice)

}
