package main

import "fmt"

func forlop() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke ", counter)
	// 	counter++
	// }

	// For dengan statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}

	// For range

	slice := []string{"Fajar", "Nugraha", "Galantixa"}
	
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// Cara cepatnya FOR range
	// names := []string{"Fajar", "Nugraha", "Galantixa"}
	for _, value  := range slice {
		// fmt.Println("Index", i , "=", value)
		fmt.Println(value)

	}
}