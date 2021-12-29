package main

import "fmt"

type Book struct {
	ID     int
	Name   string
	Author string
	Sold   int
}

func (b Book) String() string {
	return fmt.Sprintf(
		"Title:\t\t%q\n"+
			"Author:\t\t%q\n"+
			"Published:\t%v\n", b.Name, b.Author, b.Sold)
}

var Books = []Book{
	{
		ID:     1,
		Name:   "Danvanci Code",
		Author: "Dan Brown",
		Sold:   1200,
	},
	{
		ID:     2,
		Name:   "Atomic Habits",
		Author: "James Clear",
		Sold:   2000,
	},
	{
		3,
		"Power of Habit",
		"Charles Duhigg",
		3000,
	},
	{
		4,
		"Why we Sleep",
		"Mattews",
		3000,
	},
}
