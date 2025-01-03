package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"time"
)

var c, python, java bool

var i, j int = 1, 2

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var (
	tobe   bool       = false
	maxint uint64     = 1<<64 - 1
	p      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {

	var x int
	var y, javacript, cpp = 5, true, "yes"

	z := 10
	u, l, maths := true, false, "no!"

	fmt.Println("the time is", time.Now())
	fmt.Printf("Now you have %g problem.\n", math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println(add(42, 20))

	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))
	fmt.Println(x, c, python, java)

	fmt.Println(i, j, y, javacript, cpp)
	fmt.Println(z, u, l, maths)

	fmt.Printf("Type: %T value: %v\n", tobe, tobe)
	fmt.Printf("Type: %T value: %v\n", maxint, maxint)
	fmt.Printf("Type: %T value: %v\n", p, p)

	var vi int
	var vf float64
	var vb bool
	var vs string

	fmt.Printf("%v %v %v %q\n", vi, vf, vb, vs)

	var si int = 42
	var sf float64 = float64(si)
	var su uint = uint(sf)

	fmt.Println(su, sf)

	const truth = true
	fmt.Println("Go rules?", truth)

}
