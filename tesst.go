package main

import(
	"fmt"
	

)

func main(){
	var x [5]int 	//this is how you would define an array, values are initialized t zero value
	//array literal:

	var y [6]int = [6]{1,2,3,4,5,6}
	
	z:= [...]int {1,2,3,4,5} //automatic inference for what size you will need based on the number of elements in the literal


	//iterating through array

	A:= [3]int {1,2,3}

	for i,v range A {   //i will be the index and v will be bound to the value
		fmt.Printf("ind %d, val %d, i ,v)
	}
