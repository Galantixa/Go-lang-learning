package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye!!!" + name

}

func main8() {
	sayGoodBye := getGoodBye

	result := sayGoodBye("Fajar")
	fmt.Println(result)
	// hasilnya sama
	fmt.Println(getGoodBye("Fajar"))
}
