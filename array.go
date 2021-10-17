package main

import "fmt"

func main() {
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	//配列はサイズを変更できない
	var b [2]int = [2]int{100, 200}
	fmt.Println(b)
	// b = apeend(b, 300) // エラー

	//スライスはリサイズできる
	var c []int = []int{100, 200}
	c = append(c, 300)
	fmt.Println(c)
}
