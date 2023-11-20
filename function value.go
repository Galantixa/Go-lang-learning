package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye!!!" + name

}

func main() {
	// function bisa dibuat variable
	// function sayGoodBye ditampung variable contoh
	contoh := getGoodBye

	result := contoh("Fajar")
	fmt.Println(result)
	// hasilnya sama
	fmt.Println(getGoodBye("Fajar"))
}
