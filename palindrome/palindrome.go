package main 

import "fmt"

func main(){
	var num,rev int
	fmt.Println("Enter a number")
	fmt.Scan(&num)
	
	temp := num
	for temp != 0 {
		rev = rev*10 + temp%10
		temp = temp/10
	}

	if num == rev {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not Palindrome")
	}
	

}