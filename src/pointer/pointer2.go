package main

import (
	"fmt"
)

func main() {
	a := 25
	var b *int
	if b == nil {
		fmt.Println("b is", b)
		b = &a //assigning pointer to another var?
		fmt.Println("b after init is", b)
	}
}
