package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	value := "Tirta Prana Bhakti"

	var encoded string = base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println(encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println(string(decoded))
	}

}
