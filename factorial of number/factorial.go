package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}


func main() {
	var x int
	fmt.Printf("Enter a number: ")
	fmt.Scan(&x)
	fmt.Printf("Factorial of %d is %d\n", x, factorial(x))
}
