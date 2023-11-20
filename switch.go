package main

import "fmt"

func main() {
	name := "Eko"

	switch name {
	case "Fajar":
		fmt.Println("Halo Fajar")
	case "Galan":
		fmt.Println("Hallo Galan")
	default:
		fmt.Println("Kenalan Donk!")
	}

	// switch dengan short statement

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// switch length := len(name); length > 5 {
	// case true :
	// 	fmt.Println("Nama terlalu vanjang")
	// case false :
	// 	fmt.Println("Nama sudah benar")
	// }

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
