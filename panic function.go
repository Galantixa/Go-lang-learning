package main

import "fmt"

func endApp() {
	// recover adalah function yang bisa kita gunakan untuk menangkap data panic
	// dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}
	fmt.Println("App keluar")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("App Error")
	}
	fmt.Println("App berjalan")
}

func main15() {
	runApp(true) // true = panic

}
