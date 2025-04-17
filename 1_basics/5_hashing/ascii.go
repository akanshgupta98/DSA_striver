package main

import "fmt"

func main() {

	for i := 1; i <= 190; i++ {
		fmt.Printf("%d:%s %U\n", i, string(i), i)
	}
	fmt.Println()
}
