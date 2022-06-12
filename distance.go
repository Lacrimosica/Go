package main

import(
	"fmt"
)
func main(){

	var x,v,a, t float64

	fmt.Println("Do you want to calculate travelled distance?")
	
	fmt.Println("Enter initial displacement")
	fmt.Scan(&x)
	
	fmt.Println("Enter initial velocity")
	fmt.Scan(&v)
	
	fmt.Println("Enter accelaration")
	fmt.Scan(&a)

	fmt.Println("Enter time")
	fmt.Scan(&t)
	
	ffnn := GenDisplacefn(x,v,a)
	
	fmt.Println("Your Displacement is equal to : " ,ffnn(t))
}

func GenDisplacefn(xx,vv,aa float64) func(float64) float64{

	fn := func(tt float64) float64{
		return 0.5*aa*tt*tt + vv*tt + xx 
	}

	return fn
}
