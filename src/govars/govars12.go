package main

import (
	"fmt"
)

func main() {
	i := 55           // int
	j := 67.8         //int
	sum := i + int(j) // converting j into an int to pass
	fmt.Println(sum)
}
