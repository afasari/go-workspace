//Ninja Level 7
//Hands-on Exercise 2
// Create a person struct
// Create a func called "changeMe" with a *person as parameter
// 	change the value stored at the *person address
// important
// 	to deference a struct, use (*value).field
// p1.first
// (*p1).first
// 	as an exception, if the type of x is a named pointer type and (*x).f is a valid selector expression denoting a field (but not a method), x.f is shorthand for (*x).f.
// 		https://golang.org/ref/spec#Selectors
// create a value of type person
// 	print out the value
// in func main
// 	call "changeMe"
// in func main
// 	print out the value

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
