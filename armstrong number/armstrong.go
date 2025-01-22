package main

import "fmt"

func order(n int)int{
	var count int 
	for n!=0{
		count ++
		n=n/10
	}
	return count
}

func isArmstrong(n int) bool{
	temp:=n
	var sum int
	for n!=0{
		rem:=n%10
		sum+=pow(rem,order(temp))
		n=n/10
	}
	return sum==temp
}

func pow(a,b int)int{
	res:=1
	for b!=0{
		res*=a
		b--
	}
	return res
}

func main(){
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)
	if isArmstrong(n){
		fmt.Println(n," is an Armstrong number")
	}else{
		fmt.Println(n," is not an Armstrong number")
	}
	
}