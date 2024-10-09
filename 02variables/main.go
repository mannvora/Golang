package main

import "fmt"

var LoginToken string = "dfgdfgdafg" // this is a public variable and can be used by other files

func main() {
	fmt.Println("Hello from variables")

	var username string // boolean unit8 int uint16 uint32 uint 64 string // uint -> non-negative integers // int -> negative  float32 float64
	fmt.Println(username)
	fmt.Printf("username is of type %T \n", username)

	var website = "www.google.com"
	fmt.Println(website)

	users := 320 // walrus operator is not allowed globally
	fmt.Println(users)
}
