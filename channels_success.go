package main

import (
	"fmt"
)

func main() {

	c := make(chan int) //make the channel, to pull values off the channel there needs to be something to recieve it at the same time. To make that work we create a go routine to add to the channel and a print statement in the main routine to pull it off the channel

	go func() {
		c <- 42 //put the value on the channel

	}()

	fmt.Println(<-c) //pull the value off the channel
}
