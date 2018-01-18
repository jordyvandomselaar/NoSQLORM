package main

import "fmt"

func main() {
	db := NewDb()
	db.Open("./db")

	query := NewQuery(
		EqualsFilter{
			Column: "name",
			Value:  "Rotterdam",
		},
	)

	fmt.Println(db.Tables["cities"].AddQuery(query).Get())
}
