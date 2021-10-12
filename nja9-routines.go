package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//fmt.Println("vim-go")
	x := "Routine2"

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("Routine1")
		wg.Done()
	}()

	go func() {

		fmt.Println(x)
		wg.Done()
	}()

	fmt.Println("CPU's:", runtime.NumCPU())
	fmt.Println("Go Routines:", runtime.NumGoroutine())

	wg.Wait()

}
