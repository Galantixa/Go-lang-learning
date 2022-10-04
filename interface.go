package main

import "fmt"

type HasName interface {
	getName() string
}

func sayHelloo(hasName HasName) {
	fmt.Println("Hello", hasName.getName())
}

type Person struct {
	Name string
}

func (person Person) getName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) getName() string {
	return animal.Name
}
func main18() {
	var fajar Person
	fajar.Name = "galantixa"

	sayHelloo(fajar)

	cat := Animal{
		Name: "push",
	}
	sayHelloo(cat)
}
