package main

import (
	"fmt"
)

func main() {
	i := 0
	for ;i <= 10; { // init and post are omitted
		fmt.Printf("%d ", i)
		i += 2
	}
}
