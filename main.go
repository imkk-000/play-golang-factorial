package main

import "fmt"

// iterative
func factorialIterative(n float64) float64 {
	sum := float64(1)
	for c := float64(2); c <= n; c++ {
		sum *= c
	}
	return sum
}

// recursive
func factorialRecursive(n float64) float64 {
	if n == 0 {
		return 1
	}
	return n * factorialRecursive(n-1)
}

func main() {
	fmt.Println("Calculate Factorial With Iterative And Recursive")
	for c := float64(0); c <= 21; c++ {
		fmt.Printf("fac(%.0f)= int:%.0f, rec:%.0f\n", c, factorialIterative(c), factorialRecursive(c))
	}
}
