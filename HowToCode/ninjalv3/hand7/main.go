//Ninja Level 3
//Hands-on Exercise 7
// Create a program that shows the "if else if else statement" in action

package main

import "fmt"

func main() {
	x := "bond"
	if x == "bond"{
		fmt.Println("should be bond : ")
	} else if x == "james"{
		fmt.Println("should be james : ")
	} else {
		fmt.Println("should be else james or bond : ")
	}
	fmt.Println(x)
}
