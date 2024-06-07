package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var mu sync.Mutex

func main() {
	var wg sync.WaitGroup
	for i := 1; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(rand.Intn(200))
		}()
	}
	wg.Wait()
}
