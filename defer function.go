package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}


func runApplication(value int){
	defer logging() // jika tidak di defer akan panic
	fmt.Println("Run App")
	result := 10 / value
	fmt.Println("Result", result)
}

func main14() {
	runApplication(5)
}