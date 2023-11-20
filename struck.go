// struck adalah sebuah tamplate data yang digunakan untuk menggabungkan
// nol atau lebih tipe data

// data disimpan dalam field
// sederhananya struck adalah kumpulan field

// ? Membuat data struck?
// tidak bisa langsung digunakan
package main

import "fmt"

// pascal case nama depan diawali huruf kapital

type Customer struct {
	Name, Address string
	Age           int
}

//? struck function or struck method

func (customer Customer) sayHallo(Name string) {
	fmt.Println("Hello", Name, "my name is", customer.Name)
}

func main() {
	var fajar Customer
	fajar.Name = "Fajar"
	fajar.Address = "Bandung"
	fajar.Age = 19

	fajar.sayHallo("Galantixa") // memanggik struck function

	fmt.Println(fajar)
	fmt.Println(fajar.Name)
	fmt.Println(fajar.Address)
	fmt.Println(fajar.Age)

	//Struck literals

	Galantixa := Customer{
		Name:    "Galantixa",
		Address: "Bandung",
		Age:     20,
	}
	Galantixa.sayHallo("Nugraha")
	fmt.Println(Galantixa)
}
