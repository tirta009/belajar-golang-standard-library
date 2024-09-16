package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()

	data.PushBack("Tirta")
	data.PushBack("Prana")
	data.PushBack("Bhakti")

	//get ambil data pertama
	var head *list.Element = data.Front()

	//ambil data
	fmt.Println(head.Value)

	//ambil data selanjutnya
	next := head.Next()
	fmt.Println(next.Value)

	//ambil data selanjutnya
	next = next.Next()
	fmt.Println(next.Value)

	//pakai for
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
