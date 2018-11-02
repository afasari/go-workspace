//Ninja Level 3
//Hands-on Exercise 2
//Print every rune code point of the uppercase alphabet three times. Your outpush should be like this
//1.
//U+0041 'A'
//U+0041 'A'
//U+0041 'A'
//2.
//U+0042 'B'
//U+0042 'B'
//U+0042 'B'
//... throught the rest of the alphabet characters

package main

import "fmt"

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Printf("%v\n", i)
		for j:= 0; j < 3 ; j++ {
			fmt.Printf("\t%#U\n", i)

		}
	}
}
