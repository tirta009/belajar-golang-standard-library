package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	firstName := "Tirta"
	lastName := "Bhakti"

	//ini ada spasi nya
	fmt.Println("Hello '", firstName, lastName, "'")

	//dengan %s gak ada spasinya
	fmt.Printf("Hello '%s %s'\n", firstName, lastName)

}
