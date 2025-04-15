package main

import "fmt"

func printSeq(i, n int) {
	if n == 0 {
		return
	}
	fmt.Println(i)
	n--
	i++
	printSeq(i, n)
}
func main() {
	n := 10
	i := 1

	printSeq(i, n)
}
