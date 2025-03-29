package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for {
		scanner.Scan()
		line := scanner.Text()
		if len(line) == 0 {
			fmt.Println("No more input")
			break
		}
		lines = append(lines, line)

	}

	err := scanner.Err()
	if err != nil {
		fmt.Printf("Error reading from user: %s", err.Error())
		return
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}
