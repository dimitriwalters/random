package main

import (
	"fmt"
)

func test(msg string) {
	fmt.Println(msg)
}

func main() {
	for i := 0; i < 10; i++ {
		go test(fmt.Sprintf("thread%v", i))
	}

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
