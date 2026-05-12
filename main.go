package main

import (
	"fmt"
	"GO/test"
)

func main() {
	fmt.Println("hello")
	defer func(x int) {
		asd := x * x
		fmt.Println(asd)
	}(5)
	test.Status()
}

