//Fanned out to two channels and then fanned in to one channel for printing the values

package main

import (
	"fmt"
	"sync"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go send(even, odd)
	go receive(even, odd, fanin)

	for v := range fanin {
		fmt.Println(v)
	}
	fmt.Println("Program about to close")
}

//send channel
func send(even, odd chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
}

//recieve channel
func receive(even, odd <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() { //passes values onto the fanin send channel
		for v := range even {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range odd {
			fanin <- v //range will keep ranging until the channel is closed
		}
		wg.Done()
	}()

	wg.Wait() //blocks and waits for wait groups to finish executing
	close(fanin)
}
