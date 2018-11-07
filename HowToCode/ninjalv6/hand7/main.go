//Ninja Level 6
//Hands-on Exercise 7
//Assign a func to a variable. then call that func

package main

import "fmt"

func main(){
	x:=func() {
		for i:=0; i< 11; i++ {
			fmt.Println(i)
		}
	}
	x()
}