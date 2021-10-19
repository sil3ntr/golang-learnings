package main

import (
	"fmt"
	"runtime"
	"sync"
)

//trying to understand the concept https://go.dev/blog/pipelines

func main() {
	//pipeline setup - fan-out and fan-in
	in := gen(2, 3)

	//Distribute sq work over two go routines that read from in, multiple funcs can read from the same channel

	c1 := sq(in)
	c2 := sq(in)

	//consume merged output from c1 and c2
	for n := range merge(c1, c2) {
		fmt.Println(n) // 4 and 9 or 9 and 4 - depends on which routine finishes first
	}
}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums { // reads through ints in nums and sends them on the out channel
			out <- n
		}
		close(out) // closing the out send channel
	}()
	return (out) // remember to return the send channel
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n //squaring numbers being sent over close channel from func gen and passing them onto the out send channel
		}
		close(out)
	}()
	return (out)
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup //making use of a waitgroup from package sync
	out := make(chan int)

	//kicks off a go routine for each send channel
	//copies values from c to out channel is closed and then calls wg.Done

	output := func(c <-chan int) {
		for n := range c {
			out <- n //
		}
		wg.Done()
	}
	wg.Add(len(cs)) // uses length to determine number of routines for the waitgroup
	for _, c := range cs {
		go output(c) //creates the go routine for each value from cs (in our case C1 and C2 called from main)
		fmt.Println("GO Routines during Range of CS:", runtime.NumGoroutine())
	}

	//Kick off a go routine to close out once all go routines are done
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
