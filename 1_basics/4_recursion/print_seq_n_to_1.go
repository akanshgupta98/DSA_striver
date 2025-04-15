package main

import "fmt"

func printSeqRev(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	n--
	printSeqRev(n)
}
func main() {
	n := 10

	printSeqRev(n)
}
