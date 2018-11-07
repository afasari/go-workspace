//Ninja Level 4
//Hands-on Exercise 3
//Using COMPOSITE LITERAL
//	create an SLICE of TYPE int
//	assign 10 VALUES
//Range over the slice and print the values out.
//Using format printing
//	print out the TYPE of the slice

//use SLICING to create the following new slices which are then printed:
//[1 2 3 4 5]
//[6 7 8 9 10]
//[3 4 5 6 7]
//[2 3 4 5 6]

package main

import (
	"fmt"
)

func main() {
	arr := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(arr)
	fmt.Printf("%T\n",arr)
	fmt.Println(arr[:5])
	fmt.Println(arr[5:])
	fmt.Println(arr[2:7])
	fmt.Println(arr[1:6])
}
