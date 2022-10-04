package main

import "fmt"

func sumAll(numbers ...int) int{
	total := 0
	for _, value := range numbers {
		total =+ value

	}
	return total
}

func main7() {
	total := sumAll (10, 30, 20, 50, 40, 60)
	fmt.Println(total)

	slice := []int{10, 20, 30,40}
	total = sumAll(slice...)
	fmt.Println(total)
}