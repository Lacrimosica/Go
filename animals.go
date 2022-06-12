package main

import (
	"fmt"
)

type Animal struct {
	name, food, loco, spoko string
}

func (a Animal) Move() string {
	return a.loco
}

func (a Animal) Eat() string {
	return a.food
}

func (a Animal) Speak() string {
	return a.spoko
}

func (a Animal) GetName() string {
	return a.name
}

func (a Animal) DoAction(action string) string {
	switch action {
	case "move":
		return "moves by " + a.Move() + "ing"
	case "speak":
		return "says " + a.Speak()
	case "eat":
		return "eats " + a.Eat()
	default:
		return ""
	}
}

func main() {
	var (
		cow   = Animal{"cow", "grass", "walk", "moo"}
		bird  = Animal{"bird", "worms", "fly", "peep"}
		snake = Animal{"snake", "mice", "slither", "hsss"}
	)

	sl := []Animal{cow, bird, snake}

	var action, name string
	fmt.Println("Before you start typing, remember!")
	fmt.Println("You have to type everything lowercase (ex: cow move / bird eat / snake speak)")
	for {
		fmt.Print(">")

		fmt.Scanf("%s %s", &name, &action)

		for _, v := range sl {
			if name == v.GetName() {
				fmt.Println("The", v.GetName(), v.DoAction(action))
			}
		}

	}
}
