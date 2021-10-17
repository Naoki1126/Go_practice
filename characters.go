package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Hello World")
	fmt.Println("Hello" + "World")
	fmt.Println("Hello World"[0])         //72
	fmt.Println(string("Hello World"[0])) // H

	var s string = "Hello World"
	s = strings.Replace(s, "H", "X", 1) //先頭文字の置換
	fmt.Println(s)
	fmt.Println(strings.Contains(s, "World"))

	fmt.Println("Test" +
		"Test")
}
