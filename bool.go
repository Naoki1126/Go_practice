package main

import "fmt"

func main() {
	var t, f bool = true, false
	fmt.Printf("\ntype=%T value=%v %t\n", t, t, t)
	fmt.Printf("type=%T value=%v %t\n", f, f, f)

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && false)

	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || false)

	fmt.Println(!true)
	fmt.Println(!false)
}
