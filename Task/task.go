package main

import (
	"fmt"
)

// Program 1: Add and print elements into a slice (array list)
func addAndPrintElements() {
	fmt.Println("=== Program 1: Add & Print Elements ===")

	arrayList := []int{} // dynamic array (slice)

	// Adding elements
	for i := 1; i <= 5; i++ {
		arrayList = append(arrayList, i)
	}

	// Printing elements
	fmt.Println("Elements in array list:", arrayList)
}

// Program 2: Check if a number is odd or even
func checkOddEven(num int) {
	fmt.Println("\n=== Program 2: Odd or Even ===")

	if num%2 == 0 {
		fmt.Printf("%d is Even\n", num)
	} else {
		fmt.Printf("%d is Odd\n", num)
	}
}

// Program 3: Find the missing number in an array of natural numbers
// Example: Input: [1 2 4 5] → Missing: 3
func findMissingNumber(arr []int) {
	fmt.Println("\n=== Program 3: Find Missing Number ===")

	n := len(arr) + 1
	expectedSum := n * (n + 1) / 2

	actualSum := 0
	for _, v := range arr {
		actualSum += v
	}

	missing := expectedSum - actualSum
	fmt.Println("Missing Number:", missing)
}

func main() {

	// Program 1
	addAndPrintElements()

	// Program 2
	checkOddEven(11) // You can change the number

	// Program 3
	inputArray := []int{1, 2, 3, 5}
	findMissingNumber(inputArray)
}
