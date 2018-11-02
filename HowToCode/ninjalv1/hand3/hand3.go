// Exercises - Jedi Level 1
// Hands-on exercise #3
// Using the code form the previouse exercise
// 1. At the package level scope, assign the following values to the three variables
// a. for x assign 42
// b. for y assign "James Bond"
// c. for z assign true
// 2. In func main
// a. use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the returned value of TYPE string using the short declaration operator to a VARIABLE with the IDENTIFIER "s"
package main

import (
	"fmt"
)

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println("Zero Values")
	ret := fmt.Sprintf("%v\t%s\t%t", x, y, z)
	fmt.Println(ret)
}
