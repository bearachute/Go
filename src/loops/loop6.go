package main

import (
	"fmt"
)

func main() {
	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 { //mult init
		fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}
}

