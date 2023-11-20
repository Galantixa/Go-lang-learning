// go lang memiliki interface yang digunakan sebagai kontrak
// untuk  membuat error, nama interfacenya adalah error

//? membuat error
// tidak perlu manual

package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagi tidak boleh 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {
	hasil, err := Pembagian(100, 2)
	if err == nil {
		fmt.Print("Hasil ", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}
