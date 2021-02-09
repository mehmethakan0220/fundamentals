package main

import "fmt"

/*

Go programlama dilinde sınıflar yoktur. Sınıflar yerine Structlar vardır.Structlar sayesinde bir nesne oluşturulabilir ve nesneye ait özellikler oluşturabiliriz. Örnek bir struct oluşturalım.
*/

var isim string = "ali" // string türünde bir değişken tanımladık

var hayvan struct { //struct türünde bir değişken tanımladık.
	can    int
	enerji int
}

var animal struct { //Burda animal yapısını oluşturduk. Bundan bi tane örnek oluşturacağız.
	can    int
	enerji int
}

func main() {
	//hayvan değişkeninin özelliklerine manuel atama yaptık.
	hayvan.can = 100
	hayvan.enerji = 200
	fmt.Println(hayvan.can, hayvan.enerji)

	//animal yapısından bir aslan örneği oluşturalım.
	fmt.Println("animal :", animal)
	aslan := animal
	fmt.Println("aslan :", aslan)

	aslan.can = 200
	aslan.enerji = 400
	fmt.Println("aslan :", aslan)
}
