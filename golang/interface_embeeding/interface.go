package main

import "fmt"

type Animal interface {
	Speak() string
	Walk() string
}

type Dog struct {
	Name     string
	Domestic Domestic
}

type Cat struct {
	Name     string
	Domestic Domestic
}

type Domestic struct {
	isDomestic bool
	doTricks   bool
	isHungry   bool
}

func (d Dog) Speak() string {
	return "AUAU"
}

func (d Dog) Walk() string {
	return "Dog is walking"
}

func (c Cat) Speak() string {
	return "MIAU"
}

func (d *Domestic) Feed() {
	d.isHungry = false
}

func main() {
	domestic := Domestic{true, true, true}
	d := Dog{"Rex", domestic}
	c := Cat{"Felix", domestic}

	d.Domestic.Feed()

	if d.Domestic.isHungry {
		fmt.Println(d.Speak())
	}

	if c.Domestic.isHungry {
		fmt.Println(c.Speak())
	}

	fmt.Println(d.Walk())
	fmt.Println(c.Walk())
}
