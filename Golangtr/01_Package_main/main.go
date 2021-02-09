package main

func main() {
	hello() //hello.go içerisinden küçük harfle oluşturulan bir fonksiyon
	Bye()   // Bye.go içerisinde büyük harfle oluşturulan bir fonksiyon
}

/*
hello.go ve Bye.go dosya isimlerine bakilmaz büyük küçük harf farketmez. Bye.go sonradan bye.go olarak test edildi.
Burada önemli olan hello.go ve bye.go dosyalarının main paketinde olmasıdır.
Aynı pakette oldukları için fonksiyonları export etmeye gerek yoktur yani küçük harfli fonksiyon da kullanılabilir.
Export işlemi  farklı paketler arasında ileşimim için kullanılıyor.
*/
//Derlenmiş bir uygulama ilk olarak main fonksiyonuna bakar ve buradaki kodları çalıştırır.
