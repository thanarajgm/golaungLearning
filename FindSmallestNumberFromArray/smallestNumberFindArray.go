package main

import (
	"fmt"
	"sort"
)

// Method 2: Using a function
func findSmallest(arr []int) int {
	smallest := arr[0]
	for _, v := range arr {
		if v < smallest {
			smallest = v
		}
	}
	return smallest
}

func main() {
	arr := []int{10, 7, 8, 3, 5, 9, 7}

	// Method 1: Basic loop inside main
	smallest1 := arr[0]
	for _, v := range arr {
		if v < smallest1 {
			smallest1 = v
		}
	}

	// Method 2: Using a custom function
	smallest2 := findSmallest(arr)

	// Method 3 (optional): Sorting
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)
	sort.Ints(sortedArr)
	smallest3 := sortedArr[0]

	fmt.Println("Smallest (Method 1 - Loop):", smallest1)
	fmt.Println("Smallest (Method 2 - Function):", smallest2)
	fmt.Println("Smallest (Method 3 - Sort):", smallest3)
}
