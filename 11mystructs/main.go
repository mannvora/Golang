package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func main() {
	Mann := User{"Mann", "mannvora@gmail.com", 23, true}
	fmt.Println(Mann)

	fmt.Printf("%+v\n", Mann)
}
