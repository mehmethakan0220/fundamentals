package main

import (
	"fmt"
)

func main() {

	// var map_adi tanımlayıcı = tanımlayıcı{değerler}
	var m map[string]struct {
		name string
		age  int
	} = map[string]struct {
		name string
		age  int
	}{"birinci": {name: "hakan", age: 23}}

	// mapimizi define ettğimiz yerde initialize ettik. şimdi bir eleman ekleyelim

	// şuanda mapimizde keyi "birinci" olan bir struct var şimdi keyi "ikinci" olanı oluşturalım.

	m["ikinci"] = struct {
		name string
		age  int
	}{name: "mehmet", age: 25}

	fmt.Println(m)

	//şimdi burada üçüncü elemana farklı bir struct ekleyelim mesela extradan soyadı olsun.

	//m["ucuncu"]=struct{name string; surname string; age int}{}
	// burada m nesnesinin yapısına bakarak eklenecek structın farklı bir yapıda olmasına engel oluyor.

	// m nesnesinin sadece ikinci nesnesini yazdırmak istersek
	//fmt.Println(m["ikinci"])

	//birden fazla eleman ekleme. //önceki kayıtları siler

	m = map[string]struct {
		name string
		age  int
	}{
		"coklueleman1": {name: "daner", age: 25},
		"coklueleman2": {name: "irecep", age: 20},
	}

	fmt.Println(m) //önceki kayıtları sildi.
	//ekleme işi tek tek yapılmalı ve
	// m["key"] = value şeklinde olmalı.

}
