package main

import "fmt"

type employee struct {
	name    string
	surname string
	salary  int
}

func (e employee) info() { // employee tipine ait bir fonksiyon oluşturduk.
	fmt.Println("FullName:", e.name, e.surname, e.salary)
}

func main() {
	firinci := employee{
		name:    "necati",
		surname: "ergin",
		salary:  3000,
	}

	firinci.info()
}
