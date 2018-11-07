//Ninja Level 6
//Hands-on Exercise 10
//Closure is when we have "enclosed" the scope of a variable in some code block. For this hands-on exercise, create a func which encloses the scope of a variable

package main

import "fmt"

func main(){
	{
		a:= 5
		fmt.Println(a)
	}
	a:=3
	fmt.Println(a)
	f:=foo()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	g:=foo()
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
}

func foo() func() int{
	x:=0
	return func() int {
		x++
		return x
	}
}