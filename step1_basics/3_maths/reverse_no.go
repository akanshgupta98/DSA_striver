package main

import "fmt"

/*
N = 10400
o/p = 401
123
3 , 30 + 2 =32, 32*10+1 -> 321
*/
func main() {

	n := 12344
	reverseNum := 0
	for n != 0 {
		lastDigit := n % 10
		reverseNum = 10*reverseNum + lastDigit
		n /= 10
	}

	fmt.Println("Reverse is: ", reverseNum)

}
