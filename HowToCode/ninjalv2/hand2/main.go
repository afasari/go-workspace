//Ninja Level 2
//Hands-on Exercise 2
//Using the following operators, write expressions and assign their values to variables :
//g. ==
//h. <=
//i. >=
//j. !=
//k. <
//l. >
//Now Print each of the variables
package main

import (
	"fmt"
)

func main() {
	g := (42 == 42)
	h := (42 <= 42)
	i := (42 >= 42)
	j := (42 != 42)
	k := (42 < 42)
	l := (42 > 42)
	fmt.Println(g, h, i, j, k, l)
}
