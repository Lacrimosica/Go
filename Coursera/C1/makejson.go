package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

type Data struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func main() {
	var data Data

	data.Name = input("Name?")
	data.Address = input("Address?")

	output, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println("--- OUTPUT JSON ---\n", string(output))
}

func input(title string) string {
	var line string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(title + " ")
	for scanner.Scan() {
		line = strings.TrimSpace(scanner.Text())
		if utf8.RuneCountInString(line) > 0 {
			break
		}
	}
	return line
}
