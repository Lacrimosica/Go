/*
Race Condition in a program is the situation in which, the outcome of the program
depends on how the interleavings are being done.
so therefore the outcome is non-deterministic
the condition arises when two routines are working on mutual variables and their order 
is not well concurrently programmed.
*/


package main

import (
	"fmt"
	"sync"
)

func main(){

	val := 0
	var wg sync.WaitGroup

	wg.Add(2)
	//what one would expect here is that the value of val would be set to 10
	//and then printed
	//but what happens is that the printed value remains 0
	//at least my OS always schedules it in a way that the printing is done first
	//if we take the go statement from the next line, we will get 10 as output
	go dostuff(&val , &wg)

	go printStuff(&val , &wg)


	wg.Wait()
}


func dostuff(x *int , wg *sync.WaitGroup){
	
	*x = 10	
	wg.Done()
}

func printStuff(x *int , wg *sync.WaitGroup){
	fmt.Println(*x)
	wg.Done()
}


