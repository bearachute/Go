package main

import (
	"fmt"
)

func main() {
	name := "skott"
	fmt.Printf("Orig string: %s\n", string(name))
	fmt.Printf("Reversed string: ")
	for _, v := range []rune(name) {
		defer fmt.Printf("%c", v)
	}
}
