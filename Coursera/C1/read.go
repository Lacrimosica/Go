package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	type Name struct {
		fname string
		lname string
	}

	slice := make([]Name, 0, 3)
	fmt.Println("Enter the file name (.txt) : ")

	var fileName string
	fmt.Scan(&fileName)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		p := strings.Split(scanner.Text(), " ")
		var named Name
		named.fname, named.lname = p[0], p[1]
		slice = append(slice, named)
	}

	for _, v := range slice {
		fmt.Println(v.fname, v.lname)
	}

	fmt.Print(slice)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
