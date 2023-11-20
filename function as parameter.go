package main

import "fmt"

// function bisa digunakan sebagai parameter function
// filter adalah sebuah function untuk memfilter parameter name
// functon type declaration

type Filter func(string) string

func sayHelloWthFltr(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)

}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "*****"
	} else {
		return name
	}
}
func main() {
	sayHelloWthFltr("Fajar", spamFilter)
	sayHelloWthFltr("Anjing", spamFilter)
}
