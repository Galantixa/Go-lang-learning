package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke ", counter)
	// 	counter++
	// }

	// For dengan statement
	for i := 1; i <= 10; i++ {
		fmt.Println("Perulanga ke = ", i)
	}

	// For range

	slice := []string{"Fajar", "Nugraha", "Galantixa"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i], "Namanya")
	}

	// Cara cepatnya FOR range
	// names := []string{"Fajar", "Nugraha", "Galantixa"}
	// jika tidak ingin menggunakan index cukup memakai tanap _
	for _, value := range slice {
		// fmt.Println("Index", i , "=", value)
		fmt.Println(value)

	}
}
