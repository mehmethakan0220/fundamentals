package main

import "fmt"

func main() {
	// var keywordu tip belirterek oluşturma

	var isim string = "Alooov"
	var yas int = 20
	var medeniDurum bool = true

	// var keywordu ile tip belirtmeden oluşturma

	var isim2 = "neco"
	var yas2 = 83
	var medeniDurum2 = true
	//Değişken atamasında illaki değişkenin veri tipini belirtmemiz gerekmez. Yazdığımız değere göre Golang otomatik olarak veri tipini algılar.

	// en basit şekilde atama

	isim3 := "şevket"
	yas3 := 23
	medeniDurum3 := false

	/*
			Başına var eklemeden de değişken atamak mümkündür. Bu şekilde yapmak için := işaretlerini kullanırız. Aynı şekilde bu yöntemde de verinin tipi otomatik algılanır.

		Eğer değişken tanımlar iken değer kısmını boş bırakırsak yani; var yas int şeklinde yazarsak, önceki konuda da bahsettiğimiz gibi varsayılan olarak 0 değerini alır.
	*/

	fmt.Println(isim, yas, medeniDurum, isim2, yas2, medeniDurum2, isim3, yas3, medeniDurum3)
}
