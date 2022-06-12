package main


func main(){
	type struct Person {
		name string
		addr string
		phone string
	}

	var p1 Person

	
	p1.name = "joe"
	x = p1.addr

	//initializing a struct
	p1 := new(Person) //initializes the field to zero

	//using struct literal

	p1:= Person(name : "joe" , addr : "a st." , phone : "123")


}
