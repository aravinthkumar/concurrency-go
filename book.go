package main

import "fmt"

type Book struct {
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
	Book{
		Name:   "Danvanci Code",
		Author: "Dan Brown",
		Sold:   1200,
	},
	Book{
		Name:   "Atomic Habits",
		Author: "James Clear",
		Sold:   2000,
	},
	Book{
		"Power of Habit",
		"Charles Duhigg",
		3000,
	},

	Book{
		"Why we Sleep",
		"Mattews",
		3000,
	},
}
