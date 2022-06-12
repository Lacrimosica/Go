package main

import (
	"fmt"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

type Output interface {
	Eat()
	Move()
	Speak()
}

var cow = Animal{
	food:       "grass",
	locomotion: "walk",
	noise:      "moo",
}

var bird = Animal{
	food:       "worms",
	locomotion: "fly",
	noise:      "peep",
}

var snake = Animal{
	food:       "mice",
	locomotion: "slither",
	noise:      "hsss",
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

//print animal's locomotion
func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

//print animal's noise
func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	var Name string
	var Action string
	var animal Animal
	fmt.Println("Enter The Name and the action required")
	fmt.Scan(&Name, &Action)

	var n = strings.ToLower(Name)
	if n == "cow" {
		animal = cow
	} else if n == "bird" {
		animal = bird
	} else {
		animal = snake
	}

	var a = strings.ToLower(Action)
	if a == "eat" {
		animal.Eat()
	} else if a == "move" {
		animal.Move()
	} else {
		animal.Speak()
	}
}
