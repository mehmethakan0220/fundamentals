package main

import "fmt"

/*
Recursive fonksiyonlar yazdığımız fonksiyonun içinde aynı fonksiyonu kullanmamız demektir. Fonksiyonumuz tüm işlemler bittiğinde return olur.
*/

func faktoryel(sayi uint) uint {
	if sayi == 0 {
		return 1
	}
	return sayi * faktoryel(sayi-1)
}

/*
Şimdi faktoryel fonksiyonunun geri dönüş tipi int!. Ben diyorum ki int - ve + değerler alır.

return deyiminde sayi * faktoryel(sayi -1) dedik.
Bu sürekli sayi degiskenini bir eksiği ile çarpmak demek.
Normal şartlarda fonksiyon kendini her çağırdığında sayi == 0 olmadığı için direk return ifadesine geliyor ve başka bir fonksiyonu çalıştırıyor. Fonksiyon her çalıştığında sayi değerini bir azaltıp bir hesaplama yapıyor.
Bu her çalışmada sayi değeri bir azaldığı için bunun bir yerde durması lazım. Eğer if kısmını silersek.
sayi-1 ifadesi 0 altında değer almaya başlar ve son eksili değere kadar devam eder. bunun oluşmasını istemiyorsak
int yerine işaretsiz int kullanmamız gerekir yani uint.

testedildi: int yerine uint kullanım if kısmını sildiğimizde program stack overflow hatası veriyor.
benim anladığım uint tipinde olsa dahi eksi değerler alıyor ve iç içe birsürü fonksiyon çağırdığı için stak şişiyor
*/
func main() {
	fmt.Println("5 sayisinin faktöriyeli:", faktoryel(5))
}
