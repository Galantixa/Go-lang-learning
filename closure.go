// closure adalah kemampuan sebuah function berinteraksi dengan data data disekitarnya dalam scope yang sama


package main

import "fmt"

func main13() {
	counter := 0
	name := "Fajar"
	//? data atas bisa di akses 
	increment := func() {
		// name = "Nugraha" // ini akan menggantikan varibale yg di atas
		name := "Galantixa"
		fmt.Println("Increment")
		counter++
		fmt.Println(name) // hasilnya var scope yg ada di func
	//? didalam scope bawahnya
	}
	increment()
	increment()
	fmt.Println(name) // hasilnya yg ada di dalam scope var func
}