package main

import "fmt"

func main() {
	// Membuat alias untuk tipe data
	type phone string
	type maried bool
	var phoneNumber phone = "08123456788"
	var status maried = false

	fmt.Println(phoneNumber)
	fmt.Println(status)
}
