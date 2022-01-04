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
	mex := &sync.RWMutex{}
	for i := 1; i < 7; i++ {
		id := rnd.Intn(7)
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup, mex *sync.RWMutex) {
			if b, ok := getBookFromCache(id, mex); ok {
				fmt.Println("Fetched from cache")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg, mex)
		go func(id int, wg *sync.WaitGroup, mex *sync.RWMutex) {
			if b, ok := getBookFromDb(id, mex); ok {
				fmt.Println("Fetched from Database")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg, mex)
	}
	wg.Wait()
}

func getBookFromCache(id int, mex *sync.RWMutex) (Book, bool) {
	mex.RLock()
	b, ok := cache[id]
	mex.RUnlock()
	return b, ok
}

func getBookFromDb(id int, mex *sync.RWMutex) (Book, bool) {
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
