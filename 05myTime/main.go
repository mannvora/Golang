package main

import (
	"fmt"
	"time"
)

func main() {

	//go env go build
	presentTime := time.Now()

	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05"))

	// memory management -> make (non-zeroed) or new (zeroed)
	// Automatic Garbage collection -> When the GOOC goes above the threshold value  
}
