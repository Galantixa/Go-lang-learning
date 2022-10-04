package main

import "fmt"

func sayHelloWthFltr (name string, filter func (string) string) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)

}

func spamFilter (name string) string {
	if name == "Anjing" {
		return "..."
	}else {
		return name
	}
}
func main9() {
	sayHelloWthFltr("Fajar", spamFilter)
	sayHelloWthFltr("Anjing", spamFilter)
}