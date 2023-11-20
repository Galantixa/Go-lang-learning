package main

import "fmt"

func main() {
	name := "Galantixa"

	if name == "Galantixa" {
		fmt.Println("Hallo Galantixa")

	} else if name == "Joko" {
		fmt.Println("Halo Joko")
	} else {
		fmt.Println("Hai, Kenalan donks")
	}

	// if short statement

	if length := len(name); length > 12 {
		fmt.Println("terlalu panjang")

	} else {
		fmt.Println("Nama sudah benar!!!")
	}

}
