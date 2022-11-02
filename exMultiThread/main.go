package main

import (
	"fmt"
	"sync"
)

func main() {
	var compte uint32 = 0
	var wg sync.WaitGroup
	var m sync.RWMutex
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, compte *uint32, m *sync.RWMutex) {
			defer wg.Done()
			m.Lock()
			*compte++
			m.Unlock()
		}(&wg, &compte, &m)
	}
	wg.Wait()
	fmt.Println(compte)
}
