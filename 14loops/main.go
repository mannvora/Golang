package main

import "fmt"

func main() {
	fmt.Println("")

	days := []string{"Monday", "Tuesday", "Thursday", "Saturday"}

	fmt.Println(days)

	//for loop
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for i := range days {
		fmt.Printf(days[i])
	}

	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	rogueValue := 1

	for rogueValue < 10 {

		if rogueValue == 2 {
			goto lco
		}

		if rogueValue == 5 {
			break
		}
		fmt.Println(rogueValue)
		rogueValue++
	}

lco:
	fmt.Println("Jumping at xyz.com")
}
