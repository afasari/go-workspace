// Exercises - Jedi Level 1
// Hands-on exercise #1
// 1. Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the IDENTIFIER "x", "y" and "z"
// a. 42
// b. "James Bond"
// c. true

// 2. Now print the values stored in those variables using
// a. a single print statement
// b. multiple print statements
package main

import (
	"fmt"
)

func main() {
	x := 42
	y := "James Bond"
	z := true
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	// fmt.Print("\n")
	fmt.Printf("%v\t%s\t%t", x, y, z)
}
