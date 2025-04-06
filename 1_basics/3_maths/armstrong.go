package main

import (
	"fmt"
	"math"
)

/*
Armstrong No: If each digit in a no is raised to the power of no of digits, then their sum is = No.
N = 153
1^3 + 5 ^ 3 + 3 ^ 3 = 153
Then it is armstrong no.
*/
func calcPower(base, power int64) int64 {
	if power == 0 {
		return 1
	}
	if power == 1 {
		return base
	}
	temp := calcPower(base, power/2)

	if power%2 == 0 {
		return temp * temp
	}
	return base * temp * temp

}
func findDigits(n int) float64 {

	return math.Log10(float64(n)) + 1
}
func main() {

	n := int64(153)
	tmp := n
	digits := findDigits(153)
	var sum int64
	for tmp != 0 {
		result := calcPower(int64(tmp%10), int64(digits))
		sum += result
		tmp /= 10

	}

	if sum == n {
		fmt.Println("Armstrong no.")
	} else {
		fmt.Println("Not an Armstrong no.")
	}

}
