package main

import "fmt"

func main() {

	//bir dizi oluşturalım ve sınırlarının dışına çıkıp bir değer atayalım.

	dizi := [3]int{0, 1, 2}
	fmt.Println(dizi[0])
	// burada dizinin 0,1 ve 2. indisi var. 3. indise atama yapmak istersek hata verir.

	//dizi[3]=3 // yorum satırlarını kaldırırsak hata verir.

	a := []int{0, 1, 2}
	//a[3] = 3 // sliceta ekleme sadece append ile yapılıyor.
	fmt.Println(a)

	a = append(a, 4) // ayriyeten ekleme işlemi için append fonksiyonuda kullanılabilir.

	fmt.Println(a)
}
