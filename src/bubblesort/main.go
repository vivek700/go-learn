package main

import (
	"fmt"
)

func bubble_sort(slice []int) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				var temp = slice[j]
				slice[j] = slice[j+1]
				slice[j+1] = temp

			}
		}
	}

}
func binary_search(heystack []int, needle int) int {
	lo := 0
	hi := len(heystack)

	for lo < hi {
		m := lo + (hi-lo)/2
		if heystack[m] == needle {
			return m
		} else if needle > heystack[m] {
			lo = m + 1
		} else {
			hi = m
		}

	}
	return -1
}
func main() {
	slice := []int{20, 4, 80, 5, 3, 1, 6, 8, 2, 45, 89, 29, 19}
	fmt.Println(slice)
	bubble_sort(slice)
	fmt.Println(slice)
	fmt.Print(binary_search(slice, 8))
}
