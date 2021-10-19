package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	const gr = 100
	var counter int64
	var wg sync.WaitGroup
	//var m sync.Mutex

	wg.Add(gr)
	for i := 0; i < gr; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			//m.Lock()
			//c := counter
			//runtime.Gosched()
			//c++
			//counter = c
			fmt.Println(atomic.LoadInt64(&counter))
			//m.Unlock()
			wg.Done()
		}()
		fmt.Println("Go Routines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Final Counter Value: \n", counter)
	fmt.Println("Total Number of Routines: \n", runtime.NumGoroutine())
}
