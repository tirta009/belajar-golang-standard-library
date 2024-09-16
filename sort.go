package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (u UserSlice) Len() int {
	return len(u)
}

func (u UserSlice) Less(i, j int) bool {
	return u[i].Age < u[j].Age
}

func (u UserSlice) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func main() {
	users := []User{
		{"Tirta", 20},
		{"Bhakti", 30},
		{"Prana", 30},
		{"Putri", 10},
	}

	//asc
	sort.Sort(UserSlice(users))
	fmt.Println("Asc :", users)

	//desc
	sort.Sort(sort.Reverse(UserSlice(users)))
	fmt.Println("Desc :", users)

}
