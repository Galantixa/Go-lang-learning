package main

import "fmt"

func slice() {
	var months = [12]string{
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
	var slice1 = months[4:7]
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	var slice2 = append(months[2:], "Fajar")
	fmt.Println(slice2)
	fmt.Println(months)
	
	// ! NOTE!!!

	iniArray := [5]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
