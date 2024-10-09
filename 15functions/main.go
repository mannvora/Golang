package main

import "fmt"

func main() {
	var values = []int{1, 2, 4, 6, 8, 20}
	fmt.Println(findNumber(values, 5))
}

func findNumber(value []int, target int) bool {
	left := 0
	right := len(value) - 1

	for left <= right {
		mid := left + (right-left)/2

		if value[mid] == target {
			return true
		} else if value[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}
