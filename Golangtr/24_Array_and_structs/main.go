package main

import "fmt"

func main() {
	var test [2]struct {
		name    string
		surname string
		age     int
	}

	test[0] = struct {
		name    string
		surname string
		age     int
	}{"mehmet", "hakan", 22}

	test[1] = struct {
		name    string
		surname string
		age     int
	}{
		"nabi",
		"yıldız",
		29,
	}

	fmt.Println(test)

	//şimdi var ile oluşturduğumuz yerde direkt atama yapmayı deneyeceğim.

	var dizi [2]struct {
		name string
		age  int
	} = [2]struct {
		name string
		age  int
	}{{"hakan", 23}, {"nevin", 25}}

	fmt.Println(dizi)

	//şimdi aynı diziyi atama operatöründen önce tip belirtmeden yapalım.

	var dizi2 = [3]struct {
		name string
		age  int
	}{{"ece", 27}, {"aslı", 23}, {"sude", 20}}

	fmt.Println(dizi2)

	// basit var mantığı ile tanımlma yaptık. hepsi bu
	//var degisken = buraya_gelen_değere_göre_var_keywordu_degiskeni_oluşturacaktır.
	//normal var tanımı aşağıdaki gibidir
	//var degisken tip = veri
	//buradaki kural veri ile tip mutlaka birbirine uyuşmalıdır.
	//bu kuralın aslı ise şu şekildedir
	//var degisken tip = tip(veri)
	var sayi int = int(82)
	//burda dikkat edilmesi gereken atamanın sağ tarafındaki int() ifadesi
	//  () parantezleri veriye göre şekil alabilir örneğin structlarda {} ile kullanımı yaparız.
	fmt.Println(sayi)

}
