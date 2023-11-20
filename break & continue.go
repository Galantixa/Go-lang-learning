package main

import "fmt"

// break untuk menghentikan perulangan
// continou untuk melanjutkan perulangan

func main() {

	// ! BREAK
	// for i := 0; i < 10; i ++ {
	// 	if i ==  5 {
	// 		break
	// 	}
	// 	fmt.Println("Perulangan ke ", i)
	// }

	// ! CONTINUE

	for i := 0; i < 10; i++ {
		// fmt.Print("Perulangan ke ", i)
		if i%2 == 1 {
			continue
		}
		fmt.Println("Perulangan ke ", i)
	}
}
