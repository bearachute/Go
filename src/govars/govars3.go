package main

import "fmt"

func main() {
var age = 29 // type will be inferred i guess?

fmt.Println("This is different than the last because the type is inferred.")
fmt.Println("my age is", age)
}