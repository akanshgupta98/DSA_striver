package main

import "fmt"

/*
N= 5

E
D E
C D E
B C D E
A B C D E
*/
func main() {
	n := 6

	for i := 1; i <= n; i++ {
		lt := 65 + n - i
		for j := 0; j < i; j++ {
			fmt.Printf("%v ", string(lt))
			lt++
		}
		fmt.Println()
	}

}
