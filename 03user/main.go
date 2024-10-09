package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	welcome := "Welcome! Give me a value"
	// Using bufio for user input pkg.go.dev
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	// Read input and handle potential error
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred while reading input:", err)
		return
	}

	// Remove the newline character
	input = strings.TrimSpace(input)

	fmt.Println("You entered:", input)
}
