package main

import (
	"basic-golang/helper"
	"fmt"
)

func main() {
	sayHello := helper.SayHello("Galantixa")
	fmt.Println(sayHello)
	//fmt.Println(helper.version)
	fmt.Println(helper.Application)
	//fmt.Println(helper.sayGoodBye("Fajar"))
}
