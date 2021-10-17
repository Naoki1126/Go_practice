package main

import "fmt"

func main() {
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
	var result int = l[0]
	for i := 0; i < len(l); i++ {
		if i == 0 || l[i] > result {
			result = l[i]
		}
	}
	fmt.Println(result)

	m := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}
	fmt.Println(m)

	sum := 0

	for _, v := range m {
		sum += v
	}
	fmt.Println(sum)
}
