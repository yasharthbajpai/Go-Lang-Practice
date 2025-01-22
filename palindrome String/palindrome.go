package main

import "fmt"

func main(){
	var str string
	fmt.Println("Enter a string")
	fmt.Scan(&str)

	runes := []rune(str)

	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-1-i] = runes[len(runes)-1-i], runes[i]
	}

	fmt.Println("Reversed string is: ", string(runes))


}