//Ninja Level 5
//Hands-on Exercise 4
//Create and use an anonymous struct.

package main

import "fmt"

func main() {
	v1 := struct {
		test int
		sl []string
		m map[string]int
	}{
		test: 1,
		sl :[]string{"a","b"},
		m :map[string]int{
			"a" : 1,
			"b"	: 2,
		},
	}
	fmt.Println(v1)
	fmt.Println(v1.test)
	fmt.Println(v1.sl)
	fmt.Println(v1.m)
	for _, vsl := range v1.sl{
		fmt.Println(vsl)
	}
	for _, vm := range v1.m{
		fmt.Println(vm)
	}
}

