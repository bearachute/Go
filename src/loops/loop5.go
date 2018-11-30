package main

import (
	"fmt"
)

func main() {
	i := 0
	for i <= 10 { //semi are ommitted and only condition
		fmt.Printf("%d ", i)
		i += 2
	}
}
