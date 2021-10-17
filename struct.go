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
	fmt.Println(v) //{1 2 test}

	v2 := &Vertex{1, 2, "test"}
	changeVertex2(v2)
	fmt.Println(v2) //&{1000 2 test}

	/*
		v := Vertex{X: 1, Y: 2}
		fmt.Println(v)
		fmt.Println(v.X, v.Y)
		v.X = 100
		fmt.Println(v.X, v.Y)

		v2 := Vertex{X: 1}
		fmt.Println(v2)

		v3 := Vertex{1, 2, "test"}
		fmt.Println(v3.X, v3.Y, v3.S)

		v4 := Vertex{}
		fmt.Println(v4)
		fmt.Printf("%T %v\n", v4, v4)

		var v5 Vertex
		fmt.Println(v5)
		fmt.Printf("%T %v\n", v5, v5)

		v6 := new(Vertex)
		fmt.Println(v6)
		fmt.Printf("%T %v\n", v6, v6)

		v7 := &Vertex{}
		fmt.Println(v7)
		fmt.Printf("%T %v\n", v7, v7)

		s := make([]int, 0)
		fmt.Println(s)
	*/

}
