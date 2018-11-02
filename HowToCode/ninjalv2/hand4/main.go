//Ninja Level 2
//Hands-on Exercise 4
//Write a program that
//assigns an int to a variable
//prints that int in decimal, binary and hex
//shifts the bits of that int over 1 position to the left , and assign that variable
//prints that variable in decimal, binary and hex
//video 34

package main

import (
	"fmt"
)

func main() {
	var (
		a int = 42
	)
	fmt.Printf("dec:%d\tbin:%b\thex:%#X\n",a,a,a)
	a = a << 1
	fmt.Printf("dec:%d\tbin:%b\thex:%#X",a,a,a)

}
