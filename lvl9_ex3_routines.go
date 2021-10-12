package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	const gr = 100
	counter := 0
	var wg sync.WaitGroup
	var m sync.Mutex

	wg.Add(gr)
	for i := 0; i < gr; i++ {
		go func() {
			m.Lock()
			c := counter
			runtime.Gosched()
			c++
			counter = c
			fmt.Println(counter)
			m.Unlock()
			wg.Done()
		}()
		fmt.Println("Go Routines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Final Counter Value: \n", counter)
	fmt.Println("Total Number of Routines: \n", runtime.NumGoroutine())
}
