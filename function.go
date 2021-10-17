package main

import "fmt"

func add(x, y int) (int, int) {
	// fmt.Println("add function")
	// fmt.Println(x + y)
	return x + y, x - y
}

//戻り値の変数名を書いてあげることで、関数の意図が推測できる
func cal(price, item int) (result int) {
	result = price * item
	return //nakid return
}

func covert(price int) float64 {
	return float64(price)
}

func main() {
	r1, r2 := add(10, 20)
	fmt.Println(r1, r2)
	r3 := cal(100, 3)
	fmt.Println(r3)

	f := func(x int) {
		fmt.Println("innner func", x)
	}

	f(1)

	func(x int) {
		fmt.Println("innner func", x)
	}(2)
}
