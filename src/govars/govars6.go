package main

import (
"fmt"
"math"
)



func main() {
a, b := 145.8, 543.8
c := math.Min(a, b)
fmt.Println("computed at runtime vars")
fmt.Println("min value is ", c)
}