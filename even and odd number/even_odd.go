//even or odd number

package main
import "fmt"

func check(a int) int {
	if a%2==0 {
		return 1
	} else {
		return 0
	}
}

func main(){
	var a int 
	fmt.Println("Enter a number: ")
	fmt.Scanln(&a)
	if check(a)==1 {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}
	

}