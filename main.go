package main

import "fmt"

type animal interface {
	calculateFood() float32
	giveName() string
	giveWeight() float32
	typeName() string
}

type dog struct {
	weight float32
	name   string
}

func (d dog) calculateFood() float32 {
	return (d.weight / 5) * 10
}
func (d dog) giveName() string {
	return d.name
}
func (d dog) giveWeight() float32 {
	return d.weight
}
func (d dog) typeName() string {
	return "dog"
}

type cow struct {
	weight float32
	name   string
}

func (c cow) calculateFood() float32 {
	return c.weight * 25
}
func (c cow) giveName() string {
	return c.name
}
func (c cow) giveWeight() float32 {
	return c.weight
}
func (c cow) typeName() string {
	return "cow"
}

type cat struct {
	weight float32
	name   string
}

func (c cat) calculateFood() float32 {
	return c.weight * 7
}
func (c cat) giveName() string {
	return c.name
}
func (c cat) typeName() string {
	return "cat"
}
func (c cat) giveWeight() float32 {
	return c.weight
}

func main() {
	var totalFood float32
	totalFood = 0.00
	myAnimals := []animal{
		dog{35.2, "Rex"},
		cat{8.2, "Luna"},
		cow{745.2, "Daisy"}}
	for i := 0; i < len(myAnimals); i++ {
		fmt.Println("The", myAnimals[i].typeName(), "named", myAnimals[i].giveName(), "weighs", myAnimals[i].giveWeight(), "kg and needs", myAnimals[i].calculateFood(), "kg of food")
		totalFood += myAnimals[i].calculateFood()
	}
	fmt.Println(totalFood, "kg of food needed for all the animals on the farm.")
}
