//Ninja Level 5
//Hands-on Exercise 2
//Create your own type "person which will have an underlying type of "struct" so that it can be store the following data"
//	first name
//	last name
//	favorite ice cream flavors
//Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorites flavors.
//
//Store the values of type person in a map with the key of last_name.
// Access each value in the map. Print out the values, ranging over the slice.
package main

import "fmt"

type person struct {
	first_name string
	last_name string
	favorite_ice_cream_flavors []string
}
func main() {
	p1 := person{
		first_name : "afas" ,
		last_name : "ari",
		favorite_ice_cream_flavors : []string{"chocolate"},
	}
	p2 := person{
		first_name : "hailee" ,
		last_name : "steinfeld",
		favorite_ice_cream_flavors : []string{"chocolate", "vanila"},
	}
	fmt.Println(p1.first_name)
	fmt.Println(p1.last_name)
	for i1, v1:=range p1.favorite_ice_cream_flavors {
		fmt.Println(i1,v1)
	}
	fmt.Println(p2.first_name)
	fmt.Println(p2.last_name)
	for i1, v1:=range p2.favorite_ice_cream_flavors {
		fmt.Println(i1,v1)
	}
	peta := map[string]person{
		p1.last_name : p1,
		p2.last_name : p2,
	}
	for i3, v3:=range peta {
		fmt.Println("peta ",i3)
		fmt.Println(v3.first_name)
		fmt.Println(v3.last_name)

		for i4, v4:=range v3.favorite_ice_cream_flavors {
			fmt.Println(i4,v4)
		}
	}

}

