package main

import "fmt"

type katar string

func (k katar) firsLetter() (sonuc string) {
	sonuc = string(k[0])
	return
}

func main() {
	var isim katar = "Keşiş"

	fmt.Println(isim)
	fmt.Println(isim.firsLetter())
}
