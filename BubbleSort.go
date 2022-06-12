package main

import(
	"fmt"
	"strings"
	"strconv"
)


func main(){
	
	fmt.Println("enter up to 10 integers, one at a time")
	fmt.Println("enter X when done")
	
	var input string
	slice :=make([]int , 0,10)

	for i:=0 ; i<10 ; i++{
		fmt.Scan(&input)
		if strings.ToLower(input)[0] == 'x'{
			break
		}else{
			num, _ := strconv.Atoi(input)
			slice = append(slice, num)

		}
	}

	BubbleSort(slice)
	fmt.Println(slice)
}


func Swap(arr []int , i int){
	temp := arr[i]
	arr[i]=arr[i+1]
	arr[i+1] = temp

}

func BubbleSort(arr []int) {
	
	for j:=0 ; j<len(arr)-1; j++{

		for i:=0 ; i<len(arr) - j - 1  ; i++{
			
			if arr[i+1] < arr[i]{
				Swap(arr ,i)
		}
	}

	}
}
