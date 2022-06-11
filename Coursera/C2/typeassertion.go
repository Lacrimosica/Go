package main


import ("fmt")








/*interface Values
can be treated like other values:
1- assigned to vatiables
2- passed, returned

interface avlues have two componens

1. Dunamic Type: concerete type which it is assigned to
2. Dynamic Value: value of the dunamic type






*/

type Speaker interface {
	Speak()  //takes no arguments returns nothing
}


type Dog struct {name string}

func (d Dog) Speak() {
	fmt.Println(d.name)
}

func main() {
	vaar s1 Speaker
	var d1 Dog{"Brian"}
	s1 = d1  //is legal because dog type satisfies the speaker interface    
	//dynamic type here is the Dog Type
	//dynamic value is d1 which contains the name Brian

	s1.Speak()


	//interface with Nil Dynamic value

	var s2 Speaker
	var d2 *Dog
	s1= d1 //this is legal
}


func (d *Dog) Speak() {
	if d == nil {
		fmt.Println("<noise>")
	}else {
		fmt.Println(d.name)
	}
}

//the point is that that calling the methods is legal
//it's important to check for nil in the beginning of the method




//Nil interface value : in this situation you can't call the methods you want to implement on s1 if you don't do like s1 = d1




type ShapeD2 interface{    //interfaces eventually get mapped to a concrete typr
	func Area() float64
	func Parameter() float64
}


func DrawShape(s Shape2D){
	rect, ok := s.(Rectangle)	//if interface contains concrete type  rect == concrete type , ok ==true
	if ok {
		DrawRect(rect)
	}	//else rect == zer , ok = false


	tri, ok := s.(Triangle)
	if ok{
		DrawRect(tri)
	}
}

func DrawShape (s Shape3D) bool {
	switch:=sh := s.(type){ 	//sh will be the concrete value that s has
	case Rectangle:
		DrawRect(sh)
	case Triangle:
		DrawRect(sh)
	}
}


//the only way to implement an interface is to oimplement it's methods and it's satisfied automatically

type triangle {...}
func (t Triangle) Area() float64
func (t Triangle) Perimeter() float64



func FitInYard(s Shape2D) bool{
	if(s.Area() > 100 && S.Permiter() > 100 ) {
		return true
	}
	return false
}

