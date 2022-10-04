// struck adalah sebuah tamplate data yang digunakan untuk menggabungkan 
// nol atau lebih tipe data 

// data disimpan dalam field
// sederhananya struck adalah kumpulan field


//? Membuat data struck? 
// tidak bisa langsung digunakan 
package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

//? struck function or struck method

func (customer Customer) sayHallo(Name string) {
		fmt.Println("Hello", Name, "my name is", customer.Name)
}


func main17(){
	var fajar Customer
	fajar.Name = "Fajar"
	fajar.Address = "Bandung"
	fajar.Age = 19

	fajar.sayHallo("Galantixa")

	// fmt.Println(fajar)
	// fmt.Println(fajar.Name)
	// fmt.Println(fajar.Address)
	// fmt.Println(fajar.Age)

	// ? Struck literals

	// Galantixa := Customer {
	// 	Name: "Galantixa",
	// 	Address: "Bandung",
	// 	Age: 19,
	// }

	// fmt.Println(Galantixa)
}