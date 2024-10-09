package main

import "fmt"

func main() {
	var one int = 2
	fmt.Println(one)

	var ptr *int = &one
	*ptr = *ptr * 2
	fmt.Println(*ptr, ptr)

	var fruitlist = [4]string{"Tomato", "Beans", "Brinjal", "Mushrooms"}
	fmt.Println(fruitlist)
}
