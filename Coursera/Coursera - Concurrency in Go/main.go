package main

import (
	"fmt"
	"sort"
)

func sortArrayGo(slice []int, c chan []int) {
	sort.Ints(slice)
	c <- slice
}

func mySortGo(l, r []int, m chan []int) {
	resArr := make([]int, len(l)+len(r))

	for i := 0; len(l) > 0 || len(r) > 0; i++ {

		if len(l) > 0 && len(r) > 0 {
			if l[0] < r[0] {
				resArr[i] = l[0]
				l = l[1:]
			} else {
				resArr[i] = r[0]
				r = r[1:]
			}
		} else if len(l) > 0 {
			resArr[i] = l[0]
			l = l[1:]
		} else if len(r) > 0 {
			resArr[i] = r[0]
			r = r[1:]
		}
	}
	m <- resArr
}

func main() {

	arr := []int{150, 533, 58, 89, 57, 281, 1, 10}

	c := make(chan []int)
	m := make(chan []int)
	resChan := make(chan []int)

	numbersLen := len(arr)
	divideSlice := numbersLen / 4
	index := 0
	for i := 0; i < 3; i++ {
		go sortArrayGo(arr[index:index+divideSlice], c)
		index = index + divideSlice
	}
	go sortArrayGo(arr[index:], c)

	part1 := <-c
	part2 := <-c
	go mySortGo(part1, part2, m)

	part3 := <-c
	part4 := <-c
	go mySortGo(part3, part4, m)

	go mySortGo(<-m, <-m, resChan)
	fmt.Println("Result: ", <-resChan)
}
