//Ninja Level 6
//Hands-on Exercise 11
//The best way to learn is to teach. For this hands-on exercise,
//	choose one of the above exercises, or use the recursion example of factorial
//	download, install and get it running
//		https://obsproject.com/
//	record a video of YOU teaching the topic
//	upload the video on youtube
//	share the video on twitter and tag me in it so that I can see it!

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