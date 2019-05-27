package main

import (
	"fmt"
	"strconv"
)

func add(x, y int) int {
	return x + y
}

func sub(x int, y int) string {
	s := strconv.Itoa(x - y)
	return s
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(sub(42, 13))
}
