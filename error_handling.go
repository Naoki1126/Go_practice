package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./if.go")
	if err != nil {
		log.Fatalln("error", err)
	}

	defer file.Close()
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("error", err)
	}

	fmt.Println(count, string(data))

	err = os.Chdir("test")
	if err != nil {
		log.Fatalln("error", err)
	}

	if err = os.Chdir("teste"); err != nil {
		log.Fatalln("error", err)
	}
}
