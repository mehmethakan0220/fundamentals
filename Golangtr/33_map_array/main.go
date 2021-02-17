package main

import "fmt"

func main() {

	var m map[string][3]int = map[string][3]int{
		"Alinin uğurlu sayısı":   {1, 2, 3},
		"Mehmetin uğurlu sayısı": {4, 5, 6},
	}

	m["sevdanın uğurlu sayısı"] = [3]int{7, 8, 9}

	fmt.Println(m)
}
