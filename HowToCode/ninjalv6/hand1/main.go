//Ninja Level 6
//Hands-on Exercise 1
//create a func with the identifier foo that returns an int
//create a func with the identifier bar that returns an int and a string
//call both funcs
//print out their results

package main

func main() {
	foo := foo(5)
	bar1, bar2 := bar(5, "lima")
	println(foo)
	println(bar1)
	println(bar2)
}

func foo(a int)int{
	return a
}
func bar(a int, b string)(int, string){
	return a,b

}
