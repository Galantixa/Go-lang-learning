// nil hanya bisa digunakan di beberapa tipe data seperti interface, func, map, slice, pointer dan channel
package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func mai19() {
	person := NewMap("Fajar")
	fmt.Println(person)
}
