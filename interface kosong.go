// go lang bukan bahasa tipe orientasi object

package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	}else if i == 2 {
		return true
	}else {
		return "UUps"
	}
} 

func main19() {
	var data interface{} = Ups(4)
	fmt.Println(data)
}