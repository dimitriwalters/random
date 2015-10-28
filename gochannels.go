package main

import (
	"fmt"
)

func sum1() int {
	fmt.Println("sum1")
	return 5 + 3
}

func sum2() int {
	fmt.Println("sum2")
	return 12 + 1
}

func sum3() int {
	fmt.Println("sum3")
	return 4 + 7
}

func main() {
	x := make(chan int)
	y := make(chan int)
	z := make(chan int)

	go func() { x <- sum1() }()
	go func() { y <- sum2() }()
	go func() { z <- sum3() }()

	fmt.Println(fmt.Sprintf("Total: %v", <-x+<-y+<-z))
}
