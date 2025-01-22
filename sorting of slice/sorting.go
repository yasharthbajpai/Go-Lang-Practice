package main

import "fmt" 

func sort (arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i+1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
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
	fmt.Println("Sorted array is: ", sort(arr))
}
