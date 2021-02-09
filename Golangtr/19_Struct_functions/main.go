package main

import "fmt"

var employee struct { // bu employee yapısına ait bir fonksiyon oluşturamıyorum.
	name    string
	surname string
	salary  int
	info    func()
}

func main() {
	hamdi := employee

	hamdi.name = "hamdi"
	hamdi.surname = "bulan"
	hamdi.salary = 2000
	hamdi.info = func() {
		fmt.Println(hamdi.name, hamdi.surname, hamdi.salary)
	}

	hamdi.info()

}
