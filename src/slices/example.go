package main

import (
	"fmt"
	"strings"
)

func main() {
	//slices are dynamically-sized  flexible view into the elements of an array.

	names := [4]string{
		"vivek",
		"ankur",
		"tanu",
		"gungun",
	}
	x := names[0:2]
	y := names[1:3]
	fmt.Println(x, y)
	y[0] = "xxx"
	fmt.Println(x, y)
	fmt.Println(names)

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s1 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	fmt.Println(s1)

	s2 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s2)

	s2 = s2[1:4]
	fmt.Println(s2)
	printSlice(s2)

	s2 = s2[:2]
	fmt.Println(s2)
	printSlice(s2)

	s2 = s2[1:]
	fmt.Println(s2)
	printSlice(s2)

	var es []int
	fmt.Println(es, len(es), cap(es))
	if es == nil {
		fmt.Println("nil!")
	}

	ma := make([]int, 5)
	printSlice1("ma", ma)

	mb := make([]int, 0, 5)
	printSlice1("mb", mb)

	mc := mb[:2]
	printSlice1("mc", mc)

	md := mc[2:5]
	printSlice1("md", md)

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))

	}

	var appendtoslice []int
	printSlice(appendtoslice)

	appendtoslice = append(appendtoslice, 0)
	printSlice(appendtoslice)

	appendtoslice = append(appendtoslice, 1)
	printSlice(appendtoslice)

	appendtoslice = append(appendtoslice, 2, 3, 4)
	printSlice(appendtoslice)
	appendtoslice = append(appendtoslice, 5)
	printSlice(appendtoslice)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
func printSlice1(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
