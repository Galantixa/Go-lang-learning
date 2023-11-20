// indexing dimulai dari 0 - 1 - 2 .etc

package main

import "fmt"

func main() {

	// ? Array manual

	var animals [4]string
	animals[0] = "Cow"
	animals[1] = "Duck"
	animals[2] = "Cat"
	animals[3] = "Snake"

	fmt.Println(animals[0])
	fmt.Println(animals[1])
	fmt.Println(animals[2])
	fmt.Println(animals[3])
	// array index ke 3 akan error!

	// ? Array langsung saat membuat variable

	values := [3]int{
		90,
		80,
		70,
	}
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(animals))
	fmt.Println(len(values))

}
