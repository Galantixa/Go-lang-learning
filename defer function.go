package main

import "fmt"

// defer function akan dieksekusi ketika tejadi error pada function
func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	defer logging() // jika tidak di defer akan panic
	fmt.Println("Run App")
	result := 10 / value
	fmt.Println("Result", result)
}

func main() {
	runApplication(5)
}
