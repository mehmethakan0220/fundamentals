package main

import "fmt"

func main() {

	degisken := 10

	{
		degisken := 30
		fmt.Println(degisken)
		fmt.Printf("%T \t %v \n", &degisken, &degisken)
	}

	fmt.Println(degisken)
	fmt.Printf("%T \t %v \n", &degisken, &degisken)
}
