package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello")

	var fruitlist = []string{}
	fmt.Printf("%T\n", fruitlist)

	fruitlist = append(fruitlist, "apple", "banana")
	fmt.Println(fruitlist)

	//fruitlist = append(fruitlist[1:])
	fmt.Println(fruitlist)

	highScores := make([]int, 4)
	highScores[0] = 123
	highScores[1] = 223
	highScores[2] = 211
	highScores[3] = 232

	highScores = append(highScores, 245, 2466, 456)

	sort.Ints(highScores)

	fmt.Println(highScores, sort.IntsAreSorted(highScores))

	// remove value from slice based on index

	var courses = []string{"react", "rust", "golang", "node", "next"}

	fmt.Println(courses)

	var index int = 1

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
