package main

import (
	"fmt"
)

func main() {
	letter := "i"
	switch letter {
	case "a", "e", "i", "o", "u": //mult expressions
	fmt.Println("vowel")
default:
	fmt.Println("not a vowel")
}
}

