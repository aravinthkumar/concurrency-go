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
	mex := &sync.Mutex{}
	for i := 1; i < 7; i++ {
		id := rnd.Intn(7)
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup, mex *sync.Mutex) {
			if b, ok := getBookFromCache(id, mex); ok {
				fmt.Println("Fetched from cache")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg, mex)
		go func(id int, wg *sync.WaitGroup, mex *sync.Mutex) {
			if b, ok := getBookFromDb(id, mex); ok {
				fmt.Println("Fetched from Database")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg, mex)
	}
	wg.Wait()
}

func getBookFromCache(id int, mex *sync.Mutex) (Book, bool) {
	mex.Lock()
	b, ok := cache[id]
	mex.Unlock()
	return b, ok
}

func getBookFromDb(id int, mex *sync.Mutex) (Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range Books {
		if id == b.ID {
			mex.Lock()
			cache[id] = b
			mex.Unlock()
			return b, true
		}
	}
	return Book{}, false
}
