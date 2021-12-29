package main

import (
	"fmt"
	"math/rand"
	"time"
)

var cache = map[int]Book{}

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {

	for i := 0; i < 3; i++ {
		id := rnd.Intn(3) + 1
		go func(id int) {
			if b, ok := getBookFromCache(id); ok {
				fmt.Println("Fetched from cache")
				fmt.Println(b)
			}
		}(id)
		go func(id int) {
			if b, ok := getBookFromDb(id); ok {
				fmt.Println("Fetched from Database")
				fmt.Println(b)

			}
		}(id)
		fmt.Printf("Book '%v' not found", id)
		time.Sleep(1500 * time.Millisecond)
	}
	// workaround : to make the go routines run
	time.Sleep(2 * time.Second)
}

func getBookFromCache(id int) (Book, bool) {
	b, ok := cache[id]
	return b, ok
}

func getBookFromDb(id int) (Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range Books {
		if id == b.ID {
			cache[id] = b
			return b, true
		}
	}
	return Book{}, false
}
