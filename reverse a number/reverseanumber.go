package main

import "fmt"

func main() {
	var number, reverse int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	for number != 0 {
		reverse = reverse*10 + number%10
		number = number / 10
	}

	fmt.Println("Reverse of the number is:", reverse)

}