package main

import "fmt"

func main() {
	var result = 10 + 10

	fmt.Println(result)

	var a = 10
	var b = 10
	var c = a * b
	var d = a + b*c

	fmt.Println(d)

	fmt.Println(c)
	// Operasi matematika pada dirinya sendiri
	var i = 30
	i += 10 // i = i + 10
	fmt.Println(i)

	var j int8 = 11
	j++
	fmt.Println(j)
	j++
	fmt.Println(j)
	j--
	fmt.Println(j)
	j--
	fmt.Println(j)
}
