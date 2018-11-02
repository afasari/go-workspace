//Ninja Level 3
//Hands-on Exercise 4
//Create a for loop using this syntax
//	for {}
//Have it print out the year you have been alive

package main

import "fmt"

func main() {
	born := 1995;
	for {
		if (born > 2018){
			break
		}
		fmt.Println(born)
		born++
	}
}
