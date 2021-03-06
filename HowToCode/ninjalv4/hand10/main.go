//Ninja Level 4
//Hands-on Exercise 10
//Create a map with a key of TYPE string which is a person's "last_first" name,
// and a value of TYPE []string which stores their favorites things. Store three records in your map. Print out all of the values, along with their index position in the slice.
//`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
//`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
//`no_dr`, `Being evil`, `Ice cream`, `Sunsets`
//Add record  to your map. Now print the map out using the "range" loop
//Delete a record to your map. Now print the map out using the "range" loop
//
package main

import (
	"fmt"
)

func main() {
	x:= map[string][]string{
		`bond_james` : []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss` : []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr` : []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	x["afas_batiar"] = []string{`Being smart`, `Wajit`, `Heaven`}
	delete(x,`bond_james`)
	for k, v:= range x  {
		fmt.Println("This is record for ", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}

