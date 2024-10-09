package main

import "fmt"

func main() {
	var result string
	var count int = 23

	if count < 23 {
		result = "Poor"
	} else {
		result = "Good"
	}

	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("Less")
	}
}
