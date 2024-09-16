package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"John", "Paul", "George", "Ringo"}
	values := []int{100, 95, 80, 90}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))

	fmt.Println(slices.Contains(names, "Tirta"))

	//jika gak ketemu, index nya -1
	fmt.Println(slices.Index(names, "Tirta"))
	fmt.Println(slices.Index(names, "Paul"))

}
