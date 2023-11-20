package main

import "fmt"

func main() {
	// Konversi tipe data 1
	var nilai32 int32 = 3244
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)
	// Konversi tipe data 2
	var name = "Galantixa"
	var e uint8 = name[0]
	var eString string = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}
