package main

import "fmt"

/*
Dizi tanımlarken dikkat etmemiz gerekenler

var ile tanımlarken

var dizi [eleman_sayisi]int

burada [] köşeli parantezler bir dizi tanımlayacağımızı int ise verilerin integer olacağını belirtiyordu.

iki boyutlu bir dizi için de aynı kural geçerlidir.

var dizi [elemansayisi][elemanlara_ait_elemanlar]int

burada ise [elemansayisi] kısmı kaç elemanlı bir dizi olacağını belirttik.
[elemanlara_ait_elemanlar]int] burda ise her elemanın kaç int türünden kaç adet integer değer tutacağını belirttik.
*/

func main() {
	var dizi3d [2][3]int // dizi iki elemandan oluşsun her bir eleman 3 adet integer değer tutsun diyorum.
	fmt.Println(dizi3d)

	//ayni diziyi olduğu yerde initialize edelim.
	var test [2][3]int = [2][3]int{{0, 1, 2}, {3, 4, 5}}

	fmt.Println(test)

	// tanımlayıcı kullanmadan direk atama yapalım.

	var notip = [2][2]string{{"abdul", "kerim"}, {"mustafa", "kenan"}}

	fmt.Println(notip)

}
