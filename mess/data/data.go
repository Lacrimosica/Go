package data

import ("fmt")

type Point struct {
	x float64
	y float64
}

func (p *Point) InitMe(xn,yn float64){
	p.x = xn
	p.y =yn
}

func (p *Point) PrintMe(){
	fmt.Println("[", p.x ,"," , p.y ,"]")
	Println("18th april 2022 just remember for the modules to work and for running")
	Println("the data package you were stuck on this for almmost an hour, whilst being high")
}
