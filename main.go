package main

import (
	"./calc"
	"fmt"
)

func main() {
	var math calc.IMath
	math = calc.Sum{X: 5, Y: 5}
	fmt.Printf("A Soma de %d é %d\n", math, math.Calc())
}
