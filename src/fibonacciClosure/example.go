package main

import "fmt"

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		ret := a
		a, b = b, b+a
		return ret
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}

}
