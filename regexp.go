package main

import (
	"fmt"
	"regexp"
)

func main() {

	/*
		DOKUMENTASI REGEXP :
		https://github.com/google/re2/wiki/Syntax
	*/

	//buat regex nya
	var regex *regexp.Regexp = regexp.MustCompile(`e([a-z])o`)

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("edo"))
	fmt.Println(regex.MatchString("eKo"))

	fmt.Println(regex.FindAllString("eko edo egi ego e1o eto eKo", 10))
}
