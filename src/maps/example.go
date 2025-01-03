package main

import "fmt"

type Vertex struct {
	Lat, long float64
}

var m map[string]Vertex

var x = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}
var y = map[string]Vertex{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println("maps")
	fmt.Println(m)

	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	fmt.Println(m["Bell Labs"])
	fmt.Println(x)
	fmt.Println(y)

	mp := make(map[string]int)

	mp["Answer"] = 42
	fmt.Println("The value: ", mp["Answer"])
	mp["Answer"] = 48
	fmt.Println("The value: ", mp["Answer"])

	delete(mp, "Answer")
	fmt.Println("The value:", mp["Answer"])

	v, ok := mp["Answer"]
	fmt.Println("The value: ", v, "Present?", ok)

}
