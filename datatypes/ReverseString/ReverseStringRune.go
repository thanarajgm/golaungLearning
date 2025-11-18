package main

import "fmt"

func reverseString(s string) string { //Reverse Using Rune Slice
	runes := []rune(s)
	i, j := 0, len(runes)-1

	for i < j {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}

	return string(runes)
}
func reverseString2(s string) string { //Reverse Using a for-loop and String Concatenation
	reversed := ""
	for i := len(s) - 1; i >= 0; i-- {
		reversed += string(s[i])
	}
	return reversed
}
func reverseString3(s string) string { //Reverse Using Byte Slice
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
func main() {
	fmt.Println(reverseString("Java"))  // Output: avaj
	fmt.Println(reverseString2("Java")) // Output: avaj
	fmt.Println(reverseString3("Java")) // Output: avaj
}
