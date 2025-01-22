package main

import "fmt"

func freq(arr[]int) map[int]int {
	m := make(map[int]int)
	for _,v:= range arr {
		m[v]++
	}
	return m

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
	frequency := freq(arr)
	fmt.Println("Frequency of elements in the slice is: ", frequency)
}
