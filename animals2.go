package main

import (
	"fmt"
)

type Animal interface {
	Speak()
	Move()
	Eat()
	GetName() string
}

type cow struct {
	n, f, m, s string
}
type snake struct {
	n, f, m, s string
}
type bird struct {
	n, f, m, s string
}

func (s snake) Move() {
	fmt.Println(s.m)
}
func (s snake) Speak() {
	fmt.Println(s.s)
}
func (s snake) Eat() {
	fmt.Println(s.f)
}
func (s snake) GetName() string {
	return s.n
}

func (b bird) Move() {
	fmt.Println(b.m)
}
func (b bird) Speak() {
	fmt.Println(b.s)
}
func (b bird) Eat() {
	fmt.Println(b.f)
}
func (b bird) GetName() string {
	return b.n
}

func (c cow) Move() {
	fmt.Println(c.m)
}
func (c cow) Speak() {
	fmt.Println(c.s)
}
func (c cow) Eat() {
	fmt.Println(c.f)
}
func (a cow) GetName() string {
	return a.n
}

//func (a Animal) DoAction(action string) string {
//	switch action {
//	case "move":
//		return "moves by " + a.Move() + "ing"
//	case "speak":
//		return "says " + a.Speak()
//	case "eat":
//		return "eats " + a.Eat()
//	default:
//		return ""
//	}
//}

func main() {

	var input1, input2, input3 string
	fmt.Println("Before you start typing, remember!")
	fmt.Println("You have to type everything lowercase")
	fmt.Println("your choices: are\n 1. query \n 2. newanimal")
	fmt.Println("for newanimal use this typing format: \n \"newanimal <name of the new animal> <Type of the new animal> ex: newanimal milky cow\" ")
	fmt.Println("for query use this typing format:\n \"query <name of the animal> <requested info>\" ex: query ")

	var mySnake = snake{"snake", "mice", "slither", "hsss"}
	var myBird = bird{"bird", "worms", "fly", "peep"}
	var myCow = cow{"cow", "grass", "walk", "moo"}

	mySlice := []Animal{myCow, mySnake, myBird}
	//fmt.Println(mySlice)	//this line was just to make sure of the initial contents of the farm

	for {
		fmt.Print(">")

		fmt.Scanf("%s %s %s", &input1, &input2, &input3)

		switch input1 {
		case "query":
			for _, v := range mySlice {
				if input2 == v.GetName() {
					switch input3 {
					case "eat":
						v.Eat()
					case "move":
						v.Move()
					case "speak":
						v.Speak()
					}
				}
			}
		case "newanimal":
			name := input2
			typeAnimal := input3
			var NewAnimal Animal

			switch typeAnimal {
			case "cow":
				NewType := myCow
				NewType.n = name
				NewAnimal = NewType
			case "bird":
				NewType := myBird
				NewType.n = name
				NewAnimal = NewType
			case "snake":
				NewType := mySnake
				NewType.n = name
				NewAnimal = NewType
			}
			mySlice = append(mySlice, NewAnimal)
			//fmt.Println(mySlice) // this line is just for printing the content of the slice after addition to verify
			fmt.Println("Created it, a new", input3, "called", NewAnimal.GetName())
		default:
			println("I didn't get that, please repeat!")
		}
	}
}
