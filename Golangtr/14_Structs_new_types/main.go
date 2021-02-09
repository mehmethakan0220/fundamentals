package main

/*
Bu bölümde struct türünden yeni tipler türetmeyi öğreneceğiz
*/
import "fmt"

type employee struct { // struct türünde employee tipi oluşturduk
	name    string
	surname string
	id      int
}

type meslek string // string türünde meslek tipi oluşturduk.

func main() {
	temizlikPersoneli := employee{"Ayşe", "Bulan", 123456} // employe tipinden bir örnek oluşturduk
	fmt.Println(temizlikPersoneli)

	yeniMeslek := meslek("Mobilya Silici") //meslek tipinden bir örnek oluşturduk
	fmt.Println(yeniMeslek)

}
