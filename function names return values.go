package main

import "fmt"

func getFullname() (firstName string, midName string, lastname string) {
	firstName = "Fajar" 
	midName = "Nugraha"
	lastname = "Galantixa"

	return
}

func main6() {
	a, b, c := getFullname()
	

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}