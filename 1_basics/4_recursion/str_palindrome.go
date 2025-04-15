package main

import "fmt"

func checkPalin(str string, st, end int) bool {

	if st < end {
		if str[st] == str[end] {
			return checkPalin(str, st+1, end-1)
		} else {
			return false
		}

	}
	return true
}
func main() {
	str := "ABCBAB"

	fmt.Printf("The string %s is Palindrome: %v\n", str, checkPalin(str, 0, len(str)-1))
}
