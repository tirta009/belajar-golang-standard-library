package main

import (
	"fmt"
	"strconv"
)

func main() {

	//string -> boolean
	result, err := strconv.ParseBool("TRUE")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(result)
	}

	//string -> int
	resultInt, err := strconv.Atoi("1000")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(resultInt)
	}

	binary := strconv.FormatInt(999, 2)
	fmt.Println(binary)

	stringInt := strconv.Itoa(999)
	fmt.Println(stringInt)
}
