package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println("Structs")

	fmt.Println(Vertex{1, 2})

	var (
		v1 = Vertex{7, 9}
		v2 = Vertex{X: 90}
		v3 = Vertex{}
		p1 = Vertex{89, 56}
	)

	v := Vertex{5, 6}
	p := &v
	fmt.Println(v)
	fmt.Println(v.X)
	v.X = 76
	fmt.Println(v.X)

	//both are same
	fmt.Println((*p).X)
	fmt.Println(p.X)

	p.X = 1e9
	fmt.Println(v)

	fmt.Println(v1, p1, v2, v3)

}
