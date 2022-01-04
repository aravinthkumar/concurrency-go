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
		"ID:\t\t%q\n"+
			"Title:\t\t%q\n"+
			"Author:\t\t%q\n"+
			"Published:\t%v\n", b.ID, b.Name, b.Author, b.Sold)
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
		ID:     3,
		Name:   "Power of Habit",
		Author: "Charles Duhigg",
		Sold:   3000,
	},
	{
		ID:     4,
		Name:   "Why we Sleep",
		Author: "Mattews",
		Sold:   3000,
	},
	{
		ID:     5,
		Name:   "Make time",
		Author: "Mike",
		Sold:   600,
	},
	{
		ID:     6,
		Name:   "Money Matters",
		Author: "Shawn",
		Sold:   2617,
	},
	{
		ID:     7,
		Name:   "How to make friend and influence people",
		Author: "Dale Cargnie",
		Sold:   918,
	},
	{
		ID:     8,
		Name:   "Atomic Habits",
		Author: "James Clear",
		Sold:   2000,
	},
	{
		ID:     9,
		Name:   "Power of Habit",
		Author: "Charles Duhigg",
		Sold:   3000,
	},
	{
		ID:     10,
		Name:   "Why we Sleep",
		Author: "Mattews",
		Sold:   3000,
	},
	{
		ID:     11,
		Name:   "Make time",
		Author: "Mike",
		Sold:   600,
	},
	{
		ID:     12,
		Name:   "Money Matters",
		Author: "Shawn",
		Sold:   2617,
	},
	{
		ID:     13,
		Name:   "How to make friend and influence people",
		Author: "Dale Cargnie",
		Sold:   918,
	},
}
