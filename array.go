// indexing dimulai dari 0 - 1 - 2 .etc

package main

import "fmt"

func array () {

	// ? Array manual 

	var names [3] string
	names[0] = "Fajar"
	names[1] = "Nugraha"
	names[2] = "Galantixa"
	
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	// array index ke 3 akan error!

	// ? Array langsung saat membuat variable
	
	var values = [3] int {
		90,
		80,
		70,

	}
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	
	fmt.Println(len(names))
	fmt.Println(len(values))


}