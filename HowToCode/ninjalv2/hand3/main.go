//Ninja Level 2
//Hands-on Exercise 3
//Create TYPED and UNTYPED constants. Print the values of the constants

package main

import (
	"fmt"
)

const (
	a  = 42
 	b int = 42
 )
func main() {
	fmt.Printf("a :%v\t%T\nb :%v\t%T",a,a,b,b)
}
