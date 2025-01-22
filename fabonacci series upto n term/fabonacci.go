package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var x int
	fmt.Printf("Enter a number: ")
	fmt.Scan(&x)
	for i := 0; i < x; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
}

