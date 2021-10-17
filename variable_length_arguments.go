package main

import "fmt"

func foo(params ...int) {
	fmt.Println(len(params), params)

	for _, param := range params {
		fmt.Println(param)
	}

}

func main() {
	foo()
	foo(10, 20)
	foo(10, 20, 30)

	s := []int{1, 2, 3}
	fmt.Println(s)
	foo(s...)
	f := 1.11
	fmt.Println(int(f))

	s2 := []int{1, 2, 3, 4, 5, 6, 2, 3, 1}
	fmt.Println((s2[2:4]))

	m := map[string]int{"Mike": 20, "Nancy": 24, "Messi": 30}
	fmt.Printf("%T %v", m, m)
}
