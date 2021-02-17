package main

import "fmt"

func main() {

	var m = map[string]int{
		"ali": 20,
	}

	fmt.Println(m)

	//şimdi var ile tanımlayıcı kullanarak yapalım. map[string]string

	var m2 map[string]string = map[string]string{
		"ali":   "vural",
		"gamze": "onar",
		"vedat": "seven",
	}
	fmt.Println(m2)

}
