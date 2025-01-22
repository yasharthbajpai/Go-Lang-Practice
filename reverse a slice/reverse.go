package main

import "fmt"

func reverseSlice(arr []int) []int {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	return arr
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
	fmt.Println("Reversed array is: ", reverseSlice(arr))
}
