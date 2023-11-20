package main

import (
	"fmt"
)

// Tipe data absolute! Namun jika tidak digunakan tidak akan jadi masalah
func main() {
	const firstName = "Fajar"
	const lastName = "Nugraha"
	const value = 1000
	const midleNamae = "Galantixa" // ini tidak akan error

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)
}
