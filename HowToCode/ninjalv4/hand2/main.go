//Ninja Level 4
//Hands-on Exercise 2
//Using COMPOSITE LITERAL
//	create an SLICE of TYPE int
//	assign 10 VALUES
//Range over the slice and print the values out.
//Using format printing
//	print out the TYPE of the slice

package main

import (
	"fmt"
)

func main() {
	arr := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(arr)
	fmt.Printf("%T",arr)
}
