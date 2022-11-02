package main

import (
	"fmt"
	"sync"
	"time"
)

/*func main() {
	// Write your code here
	var trois uint32 = 0
	var cinq uint32 = 0
	var compte uint32 = 0
	var wg sync.WaitGroup
	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		i := i
		go func(compte *uint32, trois *uint32, cinq *uint32, wg *sync.WaitGroup) {
			defer wg.Done()
			if i%3 == 0 || i%5 == 0 {
				atomic.AddUint32(trois, 1)
			}
			atomic.AddUint32(compte, 1)
		}(&compte, &trois, &cinq, &wg)
	}
	wg.Wait()
	fmt.Print("Trois : ", trois, " Cinq : ", cinq, " Compte : ", compte)
}*/

/*func main() {
	msg := "Goroutine #"
	var m sync.Mutex
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		go func(i int, msg string, m *sync.Mutex, wg *sync.WaitGroup) {
			wg.Add(1)
			defer wg.Done()
			m.Lock()
			for j := 0; j < len(msg); j++ {
				fmt.Print(string(msg[j]))
				time.Sleep(100 * time.Millisecond)
			}
			fmt.Print(i)
			fmt.Println()
			m.Unlock()
		}(i, msg, &m, &wg)
	}
	wg.Wait()
}*/

func main() {
	var wg sync.WaitGroup
	var m sync.RWMutex
	var compte = 5

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go lire(&compte, &wg, &m)
	}
	go ecrire(&compte, &wg, &m)
	go ecrire(&compte, &wg, &m)
	wg.Wait()
}

func lire(compte *int, wg *sync.WaitGroup, m *sync.RWMutex) {
	defer wg.Done()
	m.RLock()
	fmt.Println("lire: ", *compte)
	time.Sleep(100 * time.Millisecond)
	m.RUnlock()
}

func ecrire(compte *int, wg *sync.WaitGroup, m *sync.RWMutex) {
	wg.Add(1)
	defer wg.Done()
	m.Lock()
	*compte = *compte + 1
	fmt.Println("ecrire:", *compte)
	time.Sleep(100 * time.Millisecond)
	m.Unlock()
}
