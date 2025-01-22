package main

import "fmt"

func factorsOfNumber(n int) {
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Printf("%d ", i)
		}
	}
}

func main() {
	var x int
	fmt.Printf("Enter a number: ")
	fmt.Scan(&x)
	fmt.Printf("Factors of %d are: ", x)
	factorsOfNumber(x)
}