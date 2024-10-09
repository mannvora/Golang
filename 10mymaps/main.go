package main

import "fmt"

func main() {
	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RS"] = "Rust"

	fmt.Println("List", languages)

	delete(languages, "RS")

	fmt.Println("List", languages)

	for key, value := range languages {
		fmt.Printf("For key %v the value is %v\n", key, value)
	}
}
