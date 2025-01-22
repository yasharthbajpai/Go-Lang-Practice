package main

import "fmt"

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2){
		return false
	}

	count := make(map[rune]int)

	for _, c := range s1 {
		count[c]++
	}

	for _, c := range s2 {
		count[c]--
		if count[c] < 0 {
			return false
		}
	}

	return true
}

func main(){
	var s1, s2 string
	fmt.Println("Enter two strings: ")
	fmt.Scan(&s1)
	fmt.Scan(&s2)

	if isAnagram(s1, s2) {
		fmt.Println("The strings are anagrams")
	} else {
		fmt.Println("The strings are not anagrams")
	}
}