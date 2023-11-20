package main

import "fmt"

// Interface adalah tipedata abstract, dia tidak memiliki implementasi langsung
// interface berisi defini method
// biasanya interface digunakan sebagai kontrak

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
func main() {
	var fajar Person
	fajar.Name = "Galantixa"
	sayHelloo(fajar)

	cat := Animal{Name: "push"}
	sayHelloo(cat)
}
