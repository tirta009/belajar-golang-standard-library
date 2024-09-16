package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found error")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}
	if id != "tirta" {
		return NotFoundError
	}

	//sukses

	return nil
}
func main() {
	err := GetById("tirta")
	if err != nil {

		//mengecek jenis error
		if errors.Is(err, ValidationError) {
			fmt.Println("Validation Error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("Not found Error")
		} else {
			fmt.Println("Unknown Error")
		}
	}
}
