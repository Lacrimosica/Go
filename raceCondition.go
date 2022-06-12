package main

import (
	"fmt"
	"time"
)

// Race condition is a situation whose outcome we can't predict because of bad written concurrent code.
// Race condition arises when goroutines communicate with each other.

func incrementValue(value *int) {
	*value++
}

func printValue(value *int) {
	fmt.Println(*value)
}

// We don't know what will be first: incrementing or printing into console.
func main() {
	var number = 10

	go incrementValue(&number)
	go printValue(&number)

	time.Sleep(time.Second)
}
