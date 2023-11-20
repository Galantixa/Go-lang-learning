package main

import "fmt"

// jika tipe datanya sama semua cukup deklarasikan di akhir saja
func getFullname() (firstName, midName, lastname string) {
	firstName = "Fajar"
	midName = "Nugraha"
	lastname = "Galantixa"

	return
}

func main() {
	a, b, c := getFullname()

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
