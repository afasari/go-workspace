//Ninja Level 6
//Hands-on Exercise 3
//use the "defer" keyword to show that a defered func runs after the func containing it exits.

package main

import "fmt"

func main(){
	ii:=[]int{1,2,3,4,5,6,7,8,9}

	defer foo(ii...)

	ii2:=[]int{1,2,3,4,5,6,7,8}
	bar(ii2)
}

func foo(x1 ...int) {
	total := 0
	for _, v:=range x1{
		total += v
	}
	fmt.Println(total)}
func bar(x1 []int) {
	defer func() {
		fmt.Println("defer")

	}()
	total := 0
	for _, v:=range x1{
		total += v
	}
	fmt.Println(total)
}
