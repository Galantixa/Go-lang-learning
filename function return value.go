package main

import "fmt"

func getHello(name string) string {
	return "Hello " + name
}

func main3() {
	result := getHello("Fajar") 
	fmt.Println(result)
}