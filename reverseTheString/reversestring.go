package main

import (
	"fmt"
)


func revString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	var s string
	fmt.Printf("Enter a string: ")
	fmt.Scan(&s)
	fmt.Println("Reversed string is: ", revString(s))
}