package main

import (
	"fmt"
	"strings"
)

func main(){
	var input string
	fmt.Scan(&input)
	strings.ToLower(input)
	
	if input[0] == 'i'{
		if input[len(input)-1] == 'n' {
			if strings.Contains(input, "a") {
				fmt.Println("Found!\n")
			}else{
				fmt.Println("Not Found!\n")
			}
		}
	}
}
