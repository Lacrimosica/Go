package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')

	fmt.Print("Enter your address: ")
	address, _ := reader.ReadString('\n')

	var userData = make(map[string]string)

	userData["name"] = name
	userData["address"] = address

	bytes, _ := json.Marshal(userData)
	fmt.Println(string(bytes))

}
