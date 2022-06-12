package main

import (
	"fmt"
	"bufio"
	"os"
	"sort"
	"strings"
	"strconv"
)

func main() {
	buf := bufio.NewReader(os.Stdin)	// Described: https://pkg.go.dev/bufio#NewReader

	intList := make([]int, 0, 3) 		// make a slice ints with zero length and capacity 3

	for {
		fmt.Println("Enter number for add and sort:")
		s, err := buf.ReadString('\n')		// Described: https://pkg.go.dev/bufio#Reader.ReadString
		if err != nil {
			fmt.Println(err)
			break;
		}
		s = strings.TrimSpace(s)
		if (s == "X") {
			break;
		}
		intToAdd, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
			break;
		}
		intList = append(intList, intToAdd)
		sort.Ints(intList)					// using sort std library https://pkg.go.dev/sort - specifically https://pkg.go.dev/sort#Ints
		fmt.Println(intList)
    }
}