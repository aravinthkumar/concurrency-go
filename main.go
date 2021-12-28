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
		i = rnd.Intn(3) + 1
		if b, ok := getBookFromCache(i); ok {
			fmt.Println("Fetched from cache")
			fmt.Println(b)
			continue
		}
		if b, ok := getBookFromDb(i); ok {
			fmt.Println("Fetched from Database")
			fmt.Println(b)
			continue
		}
		fmt.Printf("Book '%v' not found", i)
	}

}

func getBookFromCache(id int) (Book, bool) {
	b, ok := cache[id]
	return b, ok
}

func getBookFromDb(id int) (Book, bool) {
	for _, b := range Books {
		if id == b.ID {
			cache[id] = b
			return b, true
		}
	}
	return Book{}, false

}
