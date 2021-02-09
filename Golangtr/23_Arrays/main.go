package main

import "fmt"

func main() {

	var piyon [10]string
	//piyon adında 10 adet stringten oluşan bir dizi
	piyon[0] = "Muzaffer"
	piyon[1] = "Fuat"
	piyon[2] = "Necip"
	piyon[3] = "Murat"
	piyon[4] = "Nabi"
	piyon[5] = "Lütfü"
	piyon[6] = "Zeynal"
	piyon[7] = "Burak"
	piyon[8] = "Kasım"
	piyon[9] = "Tahsin"
	//piyon[10] = "Muzaffer"  hata verir taşma yapar.
	fmt.Println(piyon)

	// diziyi tanımladığımız yerde değer ataması yapalım bu sefer dizimiz bir uğurlu sayılarım dizisi olsun.

	var ugurlu [5]int = [5]int{1, 3, 7, 40, 63}
	var ugurlu2 = [5]int{1, 3, 7, 40, 63}
	fmt.Println(ugurlu)
	fmt.Println(ugurlu2)

	/*
		var kullanım mantığı

		var değişken tip = tipe_uygun_formatta_atama
		veya
		var değişken = herhangi_bir_tip_atama
		bu durumda uğurlu dizimizi değerlendirelim

		var uğurlu [5]int = {1, 3, 7, 40, 63}

		bu şekilde olsaydı hata verecekti ,sağ tarafın tipi belirli değildi bundan dolayı sağ tarafın da tipini net olarak belirtiyoruz.
		yani
		var ugurlu [5]int = [5]int{1,3,7,40,63}  şeklinde olmalı
		aşağıdaki örneğe bakın
	*/
	var sayi int64 = int64(29)
	fmt.Println(sayi)
	//örnekte görüldüğü gibi hata vermeden çalıştı
	//var kullanırken tipi belirtmezsek sağ taraftaki değere göre otomatik kendisi belirleyecektir.
	// uğurlu2 dizisi buna en uygun örnektir.
	//bi tane daha örnek yapalım.

	var dizi = [2]string{"sino", "selo"}
	fmt.Println(dizi)
}
