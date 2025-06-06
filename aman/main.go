package main

import (
	"fmt"
	"math"
)

func Addition(x, y int) int {
	return x + y

}

func Mul(a, b, c, d int) int {
	return a * b * c * d
}
func root(rootx int) int {
	return int(math.Sqrt(float64(rootx)))
}

func main() {
	fmt.Println("Sum:", Addition(4, 6))
	fmt.Println("Hi Aman")
	fmt.Println("sol", Mul(1, 2, 3, 4))
	fmt.Println("Sqrt:", root(44))
}
