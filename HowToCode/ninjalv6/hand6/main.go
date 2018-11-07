//Ninja Level 6
//Hands-on Exercise 6
//Build an use anonumous func

package main

import "fmt"

func main(){
	func() {
		for i:=0; i< 10; i++ {
			fmt.Println(i)
		}
	}()
}