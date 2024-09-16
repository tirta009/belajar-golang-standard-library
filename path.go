package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Dir("Hello/world.go"))
	fmt.Println(path.Base("Hello/world.go"))
	fmt.Println(path.Ext("Hello/world.go"))
	fmt.Println(path.Join("hello", "world", "main.go"))

}
