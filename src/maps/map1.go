package main

import (
	"fmt"
)

func main() {
	var personSalary map[string]int
	if personSalary == nil {
		fmt.Println("map is nil gonna make one.")
		personSalary = make(map[string]int)
	} // making a map or a hash or etc
}
