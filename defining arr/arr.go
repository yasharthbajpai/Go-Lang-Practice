package main
import "fmt"

func main() {
	var intarr [5]int
	intarr[0] = 10
	intarr[1] = 20

	fmt.Println(intarr[0])
	fmt.Println(intarr[1:3])

	fmt.Println("Length of intarr is ", len(intarr))
	fmt.Println("Capacity of intarr is ", cap(intarr))
	fmt.Println(&intarr[0])

	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)
	fmt.Println(arr2)
/////////////////////////////////////////////////////////////////

	slice1 := []int{1, 2, 3, 4, 5}
	println("Length of slice1 is ", len(slice1))
	println("Capacity of slice1 is ", cap(slice1))
	fmt.Println(slice1)
	slice1 = append(slice1, 6)
	fmt.Println(slice1)
	fmt.Printf("Length of slice1 is %d and Capacity of slice1 is %d\n", len(slice1), cap(slice1))
	slice1 = append(slice1, 4, 5, 6)	
	fmt.Println(slice1)

	slice2 := make([]int, 5, 10)
	fmt.Println(slice2)
	fmt.Printf("Length of slice2 is %d and Capacity of slice2 is %d\n", len(slice2), cap(slice2))



	
    



}