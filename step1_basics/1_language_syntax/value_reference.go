package main

import "fmt"

func Add(num1, num2 *int, result *int) {
	// result = new(int)
	*result = *num1 + *num2
}
func main() {
	// num1, num2 := 1, 2
	num1, num2 := new(int), new(int)
	var result int
	*num1 = 1
	*num2 = 2
	Add(num1, num2, &result)
	fmt.Println("Result is: ", result)

}
