package main

import "fmt"

// Tipe data slice adalah portongan dari tipedata array
// Ukuran slice bis berubah

// Pointer adalah penunjuk data pertama
// Length adalah panjang dari slice
// Capacity adalah kapasitas dari slice. length tidak boleh lebih dari capacity

// Isi data slice mengacu dari data array
func main() {
	months := [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	slice1 := months[4:7]
	fmt.Println(slice1)
	//fmt.Println(cap(slice1))
	slice2 := append(months[2:], "Fajar")
	fmt.Println(slice2)
	fmt.Println(months)

	// ! NOTE!!!

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)

	newSlice := make([]string, 3, 7)
	newSlice[0] = "Galantixa"
	newSlice[1] = "Fajar"
	newSlice[2] = "Nugraha"
	// append untuk menambahakan data ke array
	newSlice2 := append(newSlice, "Budi")
	fmt.Println(newSlice2)

	// copy slice
	//	fromSlice := slice1[:]
	//	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	//
	//	copy(fromSlice, toSlice)
	//	fmt.Println(fromSlice)
	//	fmt.Println(toSlice)
	//
	//

}
