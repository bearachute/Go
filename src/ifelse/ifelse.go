package main

import (
	"fmt"
)

func main() {
	num := 10
	if num%2 == 0 { //checks for even
		fmt.Println("the number is even")
	} else {
		fmt.Println("the number is odd")
	}
}
