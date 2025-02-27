package main

import (
	"fmt"

	"github.com/julien040/go-ternary"
)

func main() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// Ternary conditional operator
	result := ternary.If( 8%2 == 0, "either 8 or 7 are even", "either 8 or 7 are odd")
	fmt.Println(result)
}
