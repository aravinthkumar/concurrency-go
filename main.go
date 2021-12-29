package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]Book{}

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	wg := &sync.WaitGroup{}
	for i := 1; i < 7; i++ {
		id := rnd.Intn(7)
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup) {
			if b, ok := getBookFromCache(id); ok {
				fmt.Println("Fetched from cache")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg)
		go func(id int, wg *sync.WaitGroup) {
			if b, ok := getBookFromDb(id); ok {
				fmt.Println("Fetched from Database")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg)
	}
	wg.Wait()
}

func getBookFromCache(id int) (Book, bool) {
	b, ok := cache[id]
	return b, ok
}

func getBookFromDb(id int) (Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range Books {
		if id == b.ID {
			//Commenting the write operation to cache since cache is used for reading in a different place
			// which leads to race condition problems
			//cache[id] = b
			return b, true
		}
	}
	return Book{}, false
}
