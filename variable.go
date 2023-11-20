package main

import "fmt"

// ? Variable

//func main() {
//	var name string
//	name = "Galantixa"
//	fmt.Println("Nama saya", name)
//	// variable dinamis
//	name = "Fajar Nugraha"
//	fmt.Println("Nama saya", name)
//}

func main() {
	name := "Fajar Nugraha!!!"
	fmt.Println(name)
	name = "Galantixa"
	fmt.Println(name)

	// Multiple Variable

	var (
		firstName = "Fajar"
		midleName = "Galantixa"
		lastName  = "Nugraha"
	)

	fmt.Println(firstName)
	fmt.Println(midleName)
	fmt.Println(lastName)

	country := "indonesia"
	fmt.Println(country)
}
