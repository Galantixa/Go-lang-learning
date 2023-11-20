package main

import (
	"fmt"
)

func main() {
	var student1 string = "Fajar"   //type is string
	var student2 string = "Nugraha" //type is inferred
	x := 2                          //type is inferred

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
}
