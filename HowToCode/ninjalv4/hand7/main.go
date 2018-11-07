//Ninja Level 4
//Hands-on Exercise 7
//Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:
//	"James", "Bond", "Shaken, not stirred"
//	"Miss", "Moneypenny", "Helloooooo, James"
//Range over the records then range over the data in each record
package main

import (
	"fmt"
)

func main() {
	x:=[]string{"James", "Bond", "Shaken, not stirred"}
	y:=[]string{"Miss", "Moneypenny", "Helloooooo, James"}
	fmt.Println(len(x)) //lebar yang dibuat
	fmt.Println(cap(x)) //kapasitas yang disediakan
	fmt.Println(len(y)) //lebar yang dibuat
	fmt.Println(cap(y)) //kapasitas yang disediakan
	z := [][]string{x,y}
	fmt.Println(len(z)) //lebar yang dibuat
	fmt.Println(cap(z)) //kapasitas yang disediakan
	fmt.Println(z)
	for _, v1:=range z {

		for _, v2:=range v1 {
			fmt.Println(v2)
		}

	}
}

