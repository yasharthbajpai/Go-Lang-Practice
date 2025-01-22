package main

import "fmt"

func primeNumbers(n int) {
	for i :=2 ; i <= n; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}	
}

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var x int
	fmt.Printf("Enter a number: ")
	fmt.Scan(&x)
	fmt.Printf("Prime numbers till %d are: ", x)
	primeNumbers(x)
}
