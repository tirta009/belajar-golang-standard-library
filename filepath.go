package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Dir("hello\\world.go"))
	fmt.Println(filepath.Base("hello/world.go"))
	fmt.Println(filepath.Ext("hello/world.go"))
	fmt.Println(filepath.IsAbs("hello/world.go"))
	fmt.Println(filepath.IsLocal("hello/world.go"))

	//karena windows, hasilnya pakai backslash (\)
	//jika mac atau unix, pakai slash (/)
	fmt.Println(filepath.Join("hello", "world", "main.go"))

}
