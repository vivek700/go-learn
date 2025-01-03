package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//Remember: a method is just a function with a receiver argument.

// Here's Abs written as a regular function with no change in functionality.
func bbs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) vbs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)

}

func main() {
	fmt.Println("methods")

	v := Vertex{3, 4}
	fmt.Println(v.abs())
	fmt.Println(bbs(v))
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.vbs())

}
