package main

import (
	"fmt"
)

func main() {

	c := make(chan int) // general bidirectional channel expecting int
	go foo(c)           //send - call foo and provide c channel

	bar(c) //receive -

}

func foo(c chan<- int) { // here 42 will keep being added to the channel until it reaches the 100th itteration condition
	for i := 0; i < 100; i++ {
		c <- 42
	}
	close(c) //closes the channel after the loop is complete
	//range blocks the channel because unless its closed its always looking for another value to send
} //send

func bar(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
} //recieve
