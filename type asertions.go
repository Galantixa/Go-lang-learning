// type asertions merupakan kemampuan merubah tipe data menjadi tipe data yang di inginkan
// fitur ini sering sekali digunakan ketika kita bertemu dengn dara interface kosong

package main

import "fmt"

func ramdom() interface{} {
	return 1000
}

func main() {
	result := ramdom()
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	// ? lebih aman menggukana type asetions dengan switch
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown type")
	}
}
