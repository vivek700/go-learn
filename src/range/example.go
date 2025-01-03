package main

import (
	"fmt"

	"golang.org/x/tour/pic"
)

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow1 := make([]int, 10)
	for i := range pow1 {
		pow1[i] = 1 << uint(i)
	}

	for _, value := range pow1 {
		fmt.Printf("%d\n", value)
	}

	pic.Show(Pic)

}

func Pic(dx, dy int) [][]uint8 {
	slice := make([][]uint8, dy)
	for i := range slice {
		slice[i] = make([]uint8, dx)

		for x := range slice[i] {
			slice[i][x] = uint8((x * i) / 2)
		}
	}
	return slice
}
