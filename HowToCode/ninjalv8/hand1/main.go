//Ninja Level 8
//Hands-on Exercise 1

package main

import (
	"fmt"
)

type person struct {
	Name string
}

func changeMe(p *person) {
	p.Name = "Miss Moneypenny"

}
func main() {
	p1 := person{
		Name: "James Bond",
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}
