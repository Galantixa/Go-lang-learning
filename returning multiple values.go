package main

import "fmt"

func getFullName() (string, string) {
	return "Fajar", "Galantixa"

}

func main5() {
	// underscore untuk mengignore variable
	firstName, _ := getFullName()
	fmt.Println("Hallo", firstName)
}
