package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Printf("Enter the number: ")
	fmt.Scanf("%d", &n)

	
	if n < 0 {
		fmt.Printf("\n%v is not a Perfect Square!", n)
		return
	}

	
	root := int(math.Sqrt(float64(n)))

	
	if root*root == n {
		fmt.Printf("\n%v is a Perfect Square!", n)
	} else {
		fmt.Printf("\n%v is not a Perfect Square!", n)
	}
}
