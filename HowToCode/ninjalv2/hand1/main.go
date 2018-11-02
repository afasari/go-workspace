//Ninja Level 2
//Hands-on Exercise 1
//Write a program that prints a number in decimal, binary and hex
package main

import (
	"fmt"
	)

func main() {
	x:= 42
	fmt.Printf("%d\t%b\t%#X", x, x, x)
}
