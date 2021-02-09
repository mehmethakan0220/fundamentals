package main

import "fmt"

func deferic(num int) {
	defer fmt.Println("fonksiyon deferi:", num)
}

func main() {

	deferic(1)
	{
		defer fmt.Println("2.defer")
		deferic(2)
		{
			defer fmt.Println("3.defer")
			deferic(3)
			{
				defer fmt.Println("4.defer")
				deferic(4)
			}
		}
	}

	defer fmt.Println("5.defer")
	deferic(5)

}

/*
fonksiyon içindeki deferlar birbiriyle ilişkili olabilir fakat
her fonksiyonu çağırdığımızda fonksiyon görevini icra ettikten sonra defer çalışır ve yazılır.
Yani fonksiyondaki deferlar main blocktaki deferlar ile aynı stackte saklanmaz
*/
