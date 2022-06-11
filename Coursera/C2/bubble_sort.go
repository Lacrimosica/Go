package main

import "fmt"

func main() {
	fmt.Println("*** BUBBLE SORT PROGRAM ***")
	list := make([]int, 10)
	// input loop
	for i := 1; i <= 10; i++ {
		fmt.Printf("Please input an integer (%v/10): ", i)
		fmt.Scanln(&list[i-1])
	}

	BubbleSort(list)
	fmt.Println(list)
}

func BubbleSort(list []int) {
	sorted := false
	for !sorted {
		sorted = true
		for i, v := range list {
			if i == len(list)-1 {
				break
			}
			if v > list[i+1] {
				Swap(list, i)
				sorted = false
			}
		}
	}
}

func Swap(list []int, idx int) {
	cur := list[idx]
	list[idx] = list[idx+1]
	list[idx+1] = cur
}
