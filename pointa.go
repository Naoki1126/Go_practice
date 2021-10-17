package main

import "fmt"

func one(x *int) {
	*x = 1
}

func main() {
	var n int = 100
	one(&n)
	fmt.Println(n)
	fmt.Println(&n)
	fmt.Println(*&n)
	/*
		fmt.Println(n)

		fmt.Println(&n) //メモリアドレスの出力 0xc0000b2008

		var p *int = &n //ポインタ型で定義、メモリアドレスを代入する
		fmt.Println(p)  //メモリアドレスの出力

		fmt.Println(*p) //メモリアドレスの参照
	*/
}
