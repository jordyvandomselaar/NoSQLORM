package main

import "fmt"

func main() {
	db := NewDb()
	db.Open("./db")

	fmt.Println(db.Tables["cities"].Get())
}
