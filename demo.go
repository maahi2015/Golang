package main

import "fmt"

func main() {
	fmt.Println(Add(2, 1))
	fmt.Println(Subtract(2, 1))
}

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}
