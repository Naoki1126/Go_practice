package main

import (
	"fmt"
)

// func init() {
// 	fmt.Println("Init!")
// }

// func bazz() {
// 	fmt.Println("Bazz")
// }

func main() {
	//bazz()
	// fmt.Println("Hello World", time.Now())
	// fmt.Println(user.Current())
	//関数外でも定義できる
	var (
		i    int     = 1
		f64  float64 = 1.2
		s    string  = "test"
		t, f bool    = true, false
	)
	// var f bool = false
	fmt.Println(i, f64, s, t, f)
	//関数内では使える
	xi := 1
	xi = 3
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)
	fmt.Printf("\n%T\n", xf64)
	fmt.Printf("\n%T\n", xi)
}
