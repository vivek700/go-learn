package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser

	var empty interface{}
	describe(empty)

	empty = 42
	describe(empty)

	empty = "hello"
	describe(empty)

	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	a = &v
	// a = v

	fmt.Println(a.Abs())

}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
