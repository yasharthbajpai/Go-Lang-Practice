package main
import "fmt"

func main() {
	var n , sum, rem int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)
	for n > 0 {
		rem = n % 10
		sum = sum + rem
		n = n / 10
	}
	fmt.Println("Sum of digits of the number is:65", sum)

}