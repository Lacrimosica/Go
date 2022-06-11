package main

import "fmt"

func main(){
	var x float64;
	var y int;
	fmt.Println("Please enter a floating point number: ");
	fmt.Scan(&x);
	y = int(x);
	fmt.Println(y);
}