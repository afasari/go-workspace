//Ninja Level 3
//Hands-on Exercise 8
//Create a program that shows the "switch statement" with no switch expression specified

package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("ini false")
	case true:
		fmt.Println("ini true")
	}
}
