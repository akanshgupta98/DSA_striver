package main

import (
	"fmt"
	"math"
)

func countDigits(n int) (digits int) {
	//O(N)
	// for n%10 != 0 {
	// 	digits++
	// 	n /= 10
	// }
	// return
	// O(1)
	digits = int(math.Log10(float64(n))) + 1
	return

}
func main() {
	n := 7789
	fmt.Println("Digits are: ", countDigits(n))

}
