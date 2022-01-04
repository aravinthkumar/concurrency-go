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
	cacheCh := make(chan Book)
	dbCh := make(chan Book)

	for i := 1; i < 12; i++ {
		id := rnd.Intn(12) + 1
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup, mex *sync.RWMutex, ch chan<- Book) {
			if b, ok := getBookFromCache(id, mex); ok {
				ch <- b
			}
			wg.Done()
		}(id, wg, mex, cacheCh)
		go func(id int, wg *sync.WaitGroup, mex *sync.RWMutex, ch chan<- Book) {
			if b, ok := getBookFromDb(id); ok {
				mex.Lock()
				cache[id] = b
				mex.Unlock()
				ch <- b
			}
			wg.Done()
		}(id, wg, mex, dbCh)

		go func(dbCh, cacheCh <-chan Book) {
			select {
			case b := <-cacheCh:
				fmt.Println("From cache")
				fmt.Println(b)
				<-dbCh
			case b := <-dbCh:
				fmt.Println("From database")
				fmt.Println(b)
			}
		}(dbCh, cacheCh)
		time.Sleep(150 * time.Millisecond)
	}
	wg.Wait()
}

func getBookFromCache(id int, mex *sync.RWMutex) (Book, bool) {
	mex.RLock()
	b, ok := cache[id]
	mex.RUnlock()
	return b, ok
}

func getBookFromDb(id int) (Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range Books {
		if id == b.ID {

			return b, true
		}
	}
	return Book{}, false
}
