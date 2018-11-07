//Ninja Level 4
//Hands-on Exercise 1
//Using COMPOSITE LITERAL
//	create an ARRAY which holds 5 VALUES of TYPE int
//	assign VALUES to each index position.
//Range over the array and print the values out.
//Using format printing
//	print out the TYPE of the array

package main

import (
	"fmt"
)

func main() {
	arr := [5]int{1,2,3,4,5}
	fmt.Println(arr)
	fmt.Printf("%T",arr)
}
