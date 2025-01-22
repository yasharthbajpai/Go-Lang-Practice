package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Printf("Enter the number of elements: ")
	fmt.Scanf("%d", &n)

	
	if n <= 0 {
		fmt.Println("Please enter a positive number of elements.")
		return
	}

	slice := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Enter the element %d: ", i+1)
		fmt.Scanf("%d", &slice[i])
	}

	largest := slice[0]
	smallest := slice[0]

	for i := 1; i < n; i++ {
		if slice[i] > largest {
			largest = slice[i]
		}
		if slice[i] < smallest {
			smallest = slice[i]
		}
	}

	fmt.Printf("Largest element: %d\n", largest)
	fmt.Printf("Smallest element: %d\n", smallest)
}
