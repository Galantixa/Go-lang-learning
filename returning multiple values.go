package main

import "fmt"

func getFullName() (string, string) {
	return "Fajar", "Galantixa"

}

func main() {
	// underscore untuk mengignore variable
	// firtName, _ := getFullName()
	_, lastName := getFullName()

	//fmt.Println("Hallo ", firstName)
	fmt.Println("Hallo ", lastName)
}
