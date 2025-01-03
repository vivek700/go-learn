package main

import "fmt"

func main() {
	fmt.Println("loop")

	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)

	//For is Go's "while"

	//At that point you can drop the semicolons: C's while is spelled for in Go.

	sum1 := 1
	for sum1 < 1000 {
		sum1 += sum1
	}
	fmt.Println(sum1)

	//forever loop

	// for {
	// 	fmt.Println("I'm forever")
	// }

}
