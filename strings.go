package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Tirta Prana Bhakti", "Tirta"))
	fmt.Println(strings.Split("Tirta Prana Bhakti", " "))
	fmt.Println(strings.ToLower("Tirta Prana Bhakti"))
	fmt.Println(strings.ToUpper("Tirta Prana Bhakti"))
	fmt.Println(strings.Trim("         Tirta Prana Bhakti    ", " "))
	fmt.Println(strings.ReplaceAll("Tirta Prana Tirta Bhakti", "Tirta", "Bayu"))

}
