package main

import "fmt"

func main() {

	var slice []int //sadece define ettik

	var slice1 []int = []int{1, 1, 1, 1, 1} // initialize ettik ve tüm değerleri 1 yaptık. tanımlayıcı kullandık.

	var slice2 = []int{2, 2, 2, 2, 2} //initialize ettik fakat tanımlayıcı kullanmadık, tüm değerler 2

	// slicelarin tipini öğrenelim.

	fmt.Printf("type:   %T  values:   %v \n", slice, slice)
	fmt.Printf("type:   %T  values:   %v \n", slice1, slice1)
	fmt.Printf("type:   %T  values:   %v \n", slice2, slice2)

}
