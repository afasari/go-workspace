//Ninja Level 2
//Hands-on Exercise 6
// Using iota, create 4 constants for the last 4 years. Print the constant values.
//video 35
package main

import "fmt"

const (
	_     = iota
	a int = 2014 + iota
	b int = 2014 + iota
	c int = 2014 + iota
	d int = 2014 + iota
)
func main() {
	fmt.Println(a, b, c, d)

}
