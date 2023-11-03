package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func testAdd() {
	if add(1, 2) != 3 {
		panic("Test failed for add(1, 2).")
	}
	if add(1, -1) != 0 {
		panic("Test failed for add(1, -1).")
	}
}

func main() {
	testAdd()
	fmt.Println("All tests passed.")
}
