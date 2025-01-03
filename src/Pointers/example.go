package main

import "fmt"

func main() {

	fmt.Println("Pointers")
	var p *int
	fmt.Println(p)
	i := 42
	p = &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(*p)
}
