package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Fibbonacci(n int) int {
	if n <= 2 {
		return 1

	}
	last_val := 1
	ll_val := 1
	var tmp int
	for i := 3; i <= n; i++ {
		tmp = last_val + ll_val
		ll_val, last_val = last_val, tmp

	}
	return tmp
}
func main() {

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
	n, _ := strconv.Atoi(scanner.Text())
	ans := Fibbonacci(n)
	// fmt.Println(ans)
	fmt.Printf("%d fibbonacci no is: %d\n", n, ans)
}
