package main

import (
	"fmt"
)

func main() {
	arr := []int{10, 7, 8, 3, 5, 9, 7}

	// Use recover to handle panic from out-of-bound access
	safeAccess(arr, 6) // valid index
	safeAccess(arr, 7) // invalid index → throws exception
}

func safeAccess(arr []int, index int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("❌ Exception caught:", r)
		}
	}()

	fmt.Printf("Trying to access index %d...\n", index)
	value := arr[index] // may cause panic
	fmt.Println("Value =", value)
}
