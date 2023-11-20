package database

// package initializaton
// saat kita membuat sebuah package, kita bisa membuat function yang akan disakses ketika
// package kita diaskes
var connection string

func init() {
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
