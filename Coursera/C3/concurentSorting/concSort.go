package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var size int
	max := 100

	fmt.Println("input array size, i'll populate it randomly")
	_, err := fmt.Scan(&size)

	if err != nil {
		fmt.Print("something went wrong")
	}
	var array []int

	for i := 0; i < size; i++ {

		seed := rand.NewSource(time.Now().UnixNano())
		seededR := rand.New(seed)
		array = append(array, seededR.Intn(max))
	}
	fmt.Println("initial array")

	fmt.Println(array)

	subarraysize := size / 4
	remain := size % 4

	c := make(chan []int, 4)

	for j := 0; j < 4; j++ {
		if j == 3 {
			go BubbleSort(array[(j*subarraysize):((j+1)*subarraysize+remain)], c)
		} else {
			go BubbleSort(array[(j*subarraysize):((j+1)*subarraysize)], c)
		}
	}

	p1 := <-c
	p2 := <-c
	p3 := <-c
	p4 := <-c

	c = make(chan []int, 2)
	go merge(p1, p2, c)
	p5 := <-c
	go merge(p3, p4, c)
	p6 := <-c

	go merge(p5, p6, c)
	p7 := <-c

	SortedArray := p7

	fmt.Println("final Array", SortedArray)
}

func merge(s1 []int, s2 []int, c chan []int) []int {

	totalFinalLength := len(s1) + len(s2)
	var FinalArray []int
	ixa := 0
	la := len(s1)
	lb := len(s2)
	ixb := 0
	flag := true

	for flag {
		fmt.Println("len final array, tot:", len(FinalArray), totalFinalLength)
		if ixa == la && ixb == lb {
			flag = false
			c <- FinalArray
		} else if ixa < la && ixb == lb {
			FinalArray = append(FinalArray, s1[ixa])
			fmt.Println("1", FinalArray)
			ixa++
		} else if ixb < lb && ixa == la {
			FinalArray = append(FinalArray, s2[ixb])
			fmt.Println("2", FinalArray)
			ixb++
		} else {
			if s1[ixa] <= s2[ixb] {
				FinalArray = append(FinalArray, s1[ixa])
				fmt.Println("3 ", FinalArray)
				ixa++
			} else {
				FinalArray = append(FinalArray, s2[ixb])
				fmt.Println("4", FinalArray)
				ixb++
			}
		}
	}
	fmt.Println("merged slice: ", FinalArray)

	return FinalArray
}

func Swap(arr []int, i int) {
	temp := arr[i]
	arr[i] = arr[i+1]
	arr[i+1] = temp

}

func BubbleSort(arr []int, c chan []int) {
	fmt.Println("starting array:", arr)
	for j := 0; j < len(arr)-1; j++ {

		for i := 0; i < len(arr)-j-1; i++ {

			if arr[i+1] < arr[i] {
				Swap(arr, i)
			}
		}
	}
	fmt.Println(arr)
	c <- arr

}
