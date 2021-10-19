package main

import (
	"fmt"
)

func main() {

	c := make(chan int) //make the channel

	go func() {
		c <- 42 //put the value on the channel

	}()

	fmt.Println(<-c) //pull the value off the channel
}
