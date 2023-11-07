package main

import "fmt"

func main() {
	fmt.Println(Soma(20, 20))
}

// This function is responsible for sum two numbers
func Soma(a, b int) int {
	return a * b
}
