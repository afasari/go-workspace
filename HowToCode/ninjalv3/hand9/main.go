//Ninja Level 3
//Hands-on Exercise 9
//Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string
// with the IDENTIFIER "favSport"

package main

import "fmt"

func main() {
	favSport:="badminton"
	switch  favSport{
	case "badminton":
		fmt.Println("bener")
	case "futsal":
		fmt.Println("hampir")
	}
}
