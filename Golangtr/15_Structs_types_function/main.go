package main

/*
Bu bölümde struct türünde employee tipi oluşturup bu employe ait bir maas fonksiyonu hesaplatacağız
*/
import "fmt"

type employee struct {
	name    string
	surname string
	id      uint
	class   uint8
}

func (maas employee) salary() (toplam int) {
	toplam = 2000
	toplam = int(maas.class) * toplam
	return
}

func main() {
	mudur := employee{
		name:    "Ali",
		surname: "Taşkın",
		id:      732442,
		class:   5, // son level, yönetici maaşı

	}
	fmt.Println(mudur.name, "=", mudur.salary(), " lira")
}
