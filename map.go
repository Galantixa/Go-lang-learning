package main

import "fmt"

// tipe data map menggunakan key value
// data di map bisa ditambahkan sebanyak banyaknya
func main() {
	// cara cepat
	person := map[string]string{
		"name":    "Fajar",
		"address": "Bandung",
	}
	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)
	book["title"] = "belajar golang"
	book["author"] = "fajar"
	book["ups"] = "wrong"
	fmt.Println(book)

	delete(book, "ups") // menghapus data book dengan key ups dan value wrong
	fmt.Println(book)
	fmt.Println(len(book))
}
