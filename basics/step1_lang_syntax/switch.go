package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func calculateArea(choice int, arr []int) {
	var area float64
	switch choice {
	case 1:
		area = math.Pi * float64(arr[0]) * float64(arr[0])
	case 2:
		area = float64(arr[0]) * float64(arr[1])

	}
	fmt.Println("Area is: ", area)

}
func main() {

	var (
		err             error
		choice, r, l, b int
		arr             []int
	)
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	err = scanner.Err()
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
	input := scanner.Text()
	lines := strings.Split(input, " ")

	if choice, err = strconv.Atoi(lines[0]); err != nil {
		fmt.Println("Error in conversion: ", err.Error())
	}

	switch choice {
	case 1:
		r, _ = strconv.Atoi(lines[1])
		arr = []int{r}
	case 2:
		l, _ = strconv.Atoi(lines[1])
		b, _ = strconv.Atoi(lines[2])
		arr = []int{l, b}
	default:
		fmt.Println("Invalid choice")
	}

	calculateArea(choice, arr)

}
