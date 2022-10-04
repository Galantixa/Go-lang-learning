package main

import "fmt"

func typeDeclar() {
	type noID string
	type maried bool
	var myID noID = "123456789975463"
	var status maried = false

	fmt.Println(myID)
	fmt.Println(status)
}