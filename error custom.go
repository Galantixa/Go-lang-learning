package main

import "fmt"

// karena error adalah sebuah interface,jika kita ingin membuat error sendiri kita bisa membuat struck
// yang mengikuti kontrak dari interface errors

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"validation error"}
	}
	if id != "galantixa" {
		return &notFoundError{"user not found"}
	}
	// OK
	return nil
}

func main() {
	err := SaveData("galantixa", nil)
	if err != nil {
		// terjadi error
		//if validationErr, ok := err.(*validationError); ok {
		//	fmt.Println("Validation error: ", validationErr.Error())
		//} else if notFoundErr, ok := err.(*notFoundError); ok {
		//	fmt.Println("not found error: ", notFoundErr.Error())
		//} else {
		//	fmt.Println("Unknown error:", err.Error())
		//}
		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("Validation error: ", finalError.Error())
		case *notFoundError:
			fmt.Println("not found error: ", finalError.Error())
		default:
			fmt.Println("Unknown error:", finalError.Error())
		}
	} else {
		// berhasil
		fmt.Println("sukses")
	}
}
