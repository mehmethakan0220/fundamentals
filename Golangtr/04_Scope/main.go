package main

import "fmt"

func main() {

	degisken := 10
	fmt.Println(degisken)
	{
		degisken := 30
		fmt.Println(degisken)
	}
	fmt.Println(degisken)
}

/*
Yukarıda değişken isminde değişken oluşturduk. Hemen aşağısına süslü parantez oluşturduk. İçine yine değişken adında bir değişken tanımladık. Bu iki değişken aynı kod bloğunda bulunmadığı için birbirleri ile alakası olmayacaktır. Aslında ikisi de aynı değişkendir. Sadece içindeki bloğa göre farklı bir değeri vardır. Bunu anlamanın en basit yolu pointer ile bellek adresine bakmaktır. Şimdi de örneğimizin o versiyonunu görelim.
*/

/*
Şimdi sorulara geçelim
1 Bunların aynı değişken olduğunu iddia edenler var eğer bunlar aynıysa ikinci defa := kullandığımızda neden hata vermedi
2 ispatı 05_Scope_test klasöründe pointer kullanılarak test edildi.
*/
