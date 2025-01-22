package main

import "fmt"

func takeunput() (int, int, int) {
	var a, b, c int
	fmt.Println("Enter the three numbers")
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)
	return a, b, c
}

func largest(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		}
	} else if b > c {
		return b
	}
	return c
}

func main() {
	a, b, c := takeunput()
	fmt.Println("Largest of three numbers is", largest(a, b, c))
}
