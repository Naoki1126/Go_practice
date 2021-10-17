package main

import "fmt"

func main() {
	// var p *int = new(int) //adrress
	// fmt.Println(*p)
	// *p++
	// fmt.Println(*p)

	/*
		var p2 *int //nil
		fmt.Println(p2)
		*p2++
	*/

	//ポインタを返すものにはnewを使う

	s := make([]int, 0)
	fmt.Printf("%T\n", s)

	m := make(map[string]int)
	fmt.Printf("%T\n", m)

	var p *int = new(int)
	fmt.Printf("%T\n", p)

	ch := make(chan int)
	fmt.Printf("%T\n", ch)

	var st = new(struct{})
	fmt.Printf("%T\n", st)
}
