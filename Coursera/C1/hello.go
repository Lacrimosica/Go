package main

import (
	"fmt"
	"src/greetings"
)

func main(){
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
