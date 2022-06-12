package main

import (
	"fmt"
)

func main() {

	cow := make(map[string]Cow, 4)
	snake := make(map[string]Snake, 4)
	bird := make(map[string]Bird, 4)

	var animal string
	fmt.Println("Select animal: ")
	fmt.Scanln(&animal)

	var action string
	fmt.Println("Select action: ")
	fmt.Scanln(&action)

	switch animal {
	case "Cow":
		a := Cow{name: "Cow", food: "grass", locomotion: "walk", speak: "moo"}
		cow["Cow"] = a
	case "Sneak":
		a := Snake{name: "Snake", food: "mice", locomotion: "slither", speak: "hsss"}
		snake["Snake"] = a
	case "Bird":
		a := Bird{name: "Bird", food: "worms", locomotion: "fly", speak: "peep"}
		bird["Bird"] = a

	}

	switch action {
	case "Eat":
		fmt.Printf("%s eats %s\n", cow["Cow"].name, cow["Cow"].Eat())
	case "Move":
		fmt.Printf("%s moves %s\n", snake["Snake"].name, snake["Snake"].Move())
	case "Speak":
		fmt.Printf("%s speaks %s\n", bird["Bird"].name, bird["Bird"].Speak())
	}
}

type Cow struct {
	name,
	food,
	locomotion,
	speak string
}

func (animal Cow) Eat() string {
	return animal.food
}

func (animal Cow) Speak() string {
	return animal.speak
}

func (animal Cow) Move() string {
	return animal.locomotion
}

type Bird struct {
	name,
	food,
	locomotion,
	speak string
}

func (animal Bird) Eat() string {
	return animal.food
}

func (animal Bird) Speak() string {
	return animal.speak
}

func (animal Bird) Move() string {
	return animal.locomotion
}

type Snake struct {
	name,
	food,
	locomotion,
	speak string
}

func (animal Snake) Eat() string {
	return animal.food
}

func (animal Snake) Speak() string {
	return animal.speak
}

func (animal Snake) Move() string {
	return animal.locomotion
}
