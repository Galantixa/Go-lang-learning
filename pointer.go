// Pass by value
// secara default semua var itu dipassing by value, bukan reference
// artinya, jika kita mengirim sebuah variable ke dalam function, method atau variable lain
// sebenarnaya yang dikirim adalah duplikasi value nya

// Pointer
// pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama tanpa
// menduplikasi data yang sudah ada

// dengan pointer kita bisa membuat pass by reference

package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var addres1 Address = Address{"Bandung", "Jawa Barat", "Indonesia"}
	var addres2 *Address = &addres1
	// pembuktian
	var addres3 *Address = &addres1

	addres2.City = "Subang"
	// ini akan jadi data baru
	addres2 = &Address{"Jakarta", "Jakarta", "Indonesia"}

	fmt.Println(addres1)
	fmt.Println(addres2)
	fmt.Println(addres3)

	var addres4 *Address = new(Address)
	addres4.City = "Surabaya"
	fmt.Println(addres4)

}
