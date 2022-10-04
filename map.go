package main

import "fmt"

func mapp() {
	person := map[string]string{
		"name" : "Fajar",
		"addreas" : "Bandung",
	}
	person["title"] = "Programmer"

	fmt.Println(person)

	var book map[string]string = make(map[string]string)
	book["title"] = "belajar golang"
	book["author"] = "fajar"
	book["ups"] = "wrong"
	
	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))
}