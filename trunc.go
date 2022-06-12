package main

import "fmt"
import "math"
import "strconv"

func main (){
	fmt.Println("Enter a floating point number!")
	var numString string
	fmt.Scan(&numString)
	numFloat ,err := strconv.ParseFloat(numString, 64)
	_ = err
	numInt := math.Floor(numFloat)
	fmt.Println("integer part of the float is: " , numInt)
}
