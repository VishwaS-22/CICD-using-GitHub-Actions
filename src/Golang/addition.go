package main

import "fmt"

func main() {
    v1 := 11
    v2 := 11
    total := sum(v1, v2)
    fmt.Printf("The sum of %d and %d is %d\n", num1, num2, sum)
}

func sum(s, v int) int {
    return s + v
}
