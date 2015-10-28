package main

import (
	"fmt"
)

func test(msg string) {
	fmt.Println(msg)
}

func main() {
	fmt.Println("Press enter after 10 threads have executed")
	for i := 1; i <= 10; i++ {
		go test(fmt.Sprintf("thread%v", i))
	}

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
