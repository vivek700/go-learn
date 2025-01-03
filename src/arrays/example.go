package main

import "fmt"

func main() {
	fmt.Println("Arrays in go")

	var a [10]int
	fmt.Println(a)

	var s [2]string
	s[0] = "Hello"
	s[1] = "World"
	fmt.Println(s[0], s[1])
	fmt.Println(s)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	//slices are dynamically-sized  flexible view into the elements of an array.
	var slice []int = primes[1:4]
	fmt.Println(slice)

}
