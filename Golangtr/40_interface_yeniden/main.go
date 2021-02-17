package main

import "fmt"

//Actor ... fdsa
type Actor interface {
	Move(direction string)
	SaySomething(speach string)
}

//Tank ... fds
type Tank struct {
	model string
	power int
}

//Move ... fdsa
func (t Tank) Move(s string) {
	fmt.Printf("%s move to  %s \n", t.model, s)
}

//SaySomething ..
func (t Tank) SaySomething(s string) {
	fmt.Printf("%s says %s \n", t.model, s)
}

//Player ...
type Player struct {
	name string
}

//Move ..
func (p Player) Move(s string) {
	fmt.Printf("%s Move To %s \n", p.name, s)
}

//SaySomething ...
func (p Player) SaySomething(s string) {
	fmt.Printf("%s Says %s \n", p.name, s)
}

func main() {
	actors := []Actor{Tank{"Altay", 2000}, Player{"Serdar"}}

	for _, v := range actors {
		v.Move("right")
		v.SaySomething("saldırıııın")
	}

	// şimdi interface fonksiyonunu kullanarak daha okunaklı yapalım.

}
