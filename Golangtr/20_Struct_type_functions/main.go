package main

import "fmt"

type employee struct {
	name    string
	surname string
	salary  int
	info    func()
}

func main() {
	isci1 := employee{
		name:    "hasan",
		surname: "bulut",
		salary:  12000,
		info: func() {
			fmt.Println("Aktif üye")
		},
	}

	fmt.Println(isci1.name, isci1.surname)
	isci1.info()
}

// bu şekilde fonksiyon kullanmak çok yorucu olabiliyor.
