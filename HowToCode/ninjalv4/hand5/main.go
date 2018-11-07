//Ninja Level 4
//Hands-on Exercise 5
//To DELETE from a slice, we use APPEND along with SLICING. For this hands-on exercise, follow these steps:
//start with this slice
//	x := []int{42,43,44,45,46,47,48,49,50,51}
//use APPEND & SLICING to get the values here which you should then print:
//	[42,43,44,48,49,50,51]

package main

import (
	"fmt"
)

func main() {
	x:=[]int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)
	x = append(x[:3], x[6:]...)
	fmt.Println(x)
}
