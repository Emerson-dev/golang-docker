package main

import "fmt"

func main() {
	total := Sum(5, 5)
	fmt.Printf("A soma de 5 + 5 = %d\n", total)
}

func Sum(x int, y int) int {
	return x + y
}