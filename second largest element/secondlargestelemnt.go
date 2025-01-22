package main

import "fmt"

func seclargElem(arr []int) int {
	first := arr[0]
	second := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] > first {
			second = first
			first = arr[i]
		} else if arr[i] > second && arr[i] != first {
			second = arr[i]
		}
	}
	return second
}

func main() {
	var n int
	fmt.Printf("Enter the number of elements: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	fmt.Println("Enter the elements: ")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Printf("Second largest element is: %d\n", seclargElem(arr))
}
