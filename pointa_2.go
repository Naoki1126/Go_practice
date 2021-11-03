package main

import "fmt"

type Vertex struct {
	X, Y int
	S    string
}

func changeVertex(v Vertex) {
	v.X = 1000

}
func changeVertex2(v *Vertex) {
	v.X = 1000

}

func main() {
	v := Vertex{1, 2, "test"}
	changeVertex(v)
	fmt.Println(v)

	v2 := &Vertex{1, 2, "test"}
	changeVertex2(v2)
	fmt.Println(v2)
	// changeVertex2()
	// 	var n int = 100
	// 	fmt.Println(n)

	// 	fmt.Println(&n)

	// 	var p *int = &n
	// 	fmt.Println(p)

	// 	fmt.Println(*p)

	// 	var v5 Vertex
	// 	fmt.Println(v5)

	// 	v6 := new(Vertex)
	// 	fmt.Println(v6)

	// 	v7 := &Vertex{}
	// 	fmt.rintln(v7)
}
