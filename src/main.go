package main

import "fmt"

func main() {
	fmt.Println(add1(1, 2))
	fmt.Println(add2(1.2, 2.3))
	fmt.Println(add(1, 9))
	fmt.Println(add(1.1, 8.9))
}

func add1(a, b int) int {
	return a + b
}

func add2(x, y float64) float64 {
	return x + y
}

type addable interface {
	int | int64 | float64
}

func add[T addable](x, y T) T {
	return x + y
}
