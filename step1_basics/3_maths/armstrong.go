package main

/*
Armstrong No: If each digit in a no is raised to the power of no of digits, then their sum is = No.
N = 153
1^3 + 5 ^ 3 + 3 ^ 3 = 153
Then it is armstrong no.
*/
func power(base, power int64) int64 {
	var pr int64
	var i int64
	for i = 0; i < power; i++ {
		pr = pr * base
	}
	return pr

}
func main() {

}
