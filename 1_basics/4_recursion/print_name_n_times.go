package main

import "fmt"

func printName(n int) {
	if n == 0 {
		return
	}
	fmt.Println("Akansh")
	n--
	printName(n)
}
func main() {
	n := 10

	printName(n)
}
