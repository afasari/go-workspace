//Ninja Level 6
//Hands-on Exercise 9
//A "callback" is when we pass a func into a func as an argument. for this exercise,
//	pass a func into a func as an argument

package main

import "fmt"

func main(){
	g := func(x1 []int) int {
		if len(x1) == 0 {
			return 0
		} else if len(x1) == 1 {
			return x1[0]
		}
		return x1[0] + x1[len(x1)-1]
	}
	x := foo(g, []int{1,2,3,4,5})
	fmt.Println(x)
}
func foo(f func(x1 []int) int, ii []int) int{
	n := f(ii)
	n++
	return n
}