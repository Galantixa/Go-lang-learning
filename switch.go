package main

import "fmt"

func swithc() {
	name := "Ono"

	switch name {
	case "Fajar" :
		fmt.Println("Hello, Fajar")
	case "Joko" :
		fmt.Println("Hello, Joko")
	default:
		fmt.Println("Kenalan donkk")
	}
	
	// switch dengan short statement 

	switch length := len(name); length > 5 {
	case true :
		fmt.Println("Nama terlalu panjang")
	case false :
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
