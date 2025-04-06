package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
Check if N is palindrome or not.
N = 121.

1. If a reverse of a no is same as n, then it is palindrome.
*/

func reverse(n int) (rev int) {
	for n != 0 {
		lastDigit := n % 10
		rev = 10*rev + lastDigit
		n /= 10
	}
	return
}
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	// n := 121
	n1 := reverse(n)
	if n == n1 {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not Palindrome")
	}

}
