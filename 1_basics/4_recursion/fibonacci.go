package main

import "fmt"

func NFib(n int) int {

	if n == 0 || n == 1 {
		// fmt.Printf("%d ", n)
		return n
	}
	l := NFib(n - 2)
	// if l == 0 || l == 1 {
	// 	fmt.Printf("%d ", l)
	// }
	l2 := NFib(n - 1)
	// if l2 == 0 || l2 == 1 {
	// 	fmt.Printf("%d ", l2)
	// }
	// fmt.Printf("%d ", l+l2)
	return l + l2
	// 5 -> 4 + 3.
	//  -> 3 + 2
	//  -> 2 + 1
	//  -> 1 + 0

	// fmt.Printf("%d ", tmp)
	// if n <

	// return tmp

}
func main() {
	n := 5
	for i := range n {
		fmt.Printf("%d ", NFib(i))
	}
	// last := 1
	// last2 := 0
	// fmt.Printf("%d ", last2)
	// fmt.Printf("%d ", last)

	// for i := 2; i <= n; i++ {
	// 	fmt.Printf("%d ", last+last2)
	// 	last, last2 = last+last2, last

	// }
	// fib(n,0)
	// fmt.Println(i)
	// fib(n,i+1)
	// fmt.Printf("\n%d th Fibbonacci no is: %d\n", n, NFib(n))
}
