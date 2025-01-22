package main

import (
    "fmt"
    "strings"
)

func main() {
    var str string
    vo := 0
    fmt.Printf("Enter a string: ")
    fmt.Scanln(&str) 
    str = strings.ToLower(str) 

    for i := 0; i < len(str); i++ { 
        if str[i] == 'a' || str[i] == 'e' || str[i] == 'i' || str[i] == 'o' || str[i] == 'u' {
            vo++ 
        }
    }

    fmt.Printf("Number of vowels in the string are: %d\n", vo)
}
