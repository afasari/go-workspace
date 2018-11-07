//Ninja Level 6
//Hands-on Exercise 5
//Create a type SQUARE
//create a type CIRCLE
//attach a method to each that calculates AREA and returns it
//	circle area = phi r kuadrat
//	square area = p kali l
//create type SHAPE which defines an interface as anything which has the AREA method
//create a func INFO which takes type shape and then prints the area
//create a value of type square
//create a value of type circle
//use func info to print the area of square
//use func info to print the area of circle

package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (p circle) area() float64{
	return math.Pi * p.radius * p.radius
}

func (p square) area() float64{
	return p.length * p.length
}

type shape interface {
	area() float64
}

func info(s shape)  {
	x:=s.area()
	fmt.Println(x)
}
func main(){
	c:=circle{
		radius:123,
	}
	s:=square{
		length:123,
	}
	info(c)
	info(s)
}