package main

import (
	"fmt"
)

func main() {

	c := make(chan<- int, 2) //here we have a buffer of 2, this is directional, and only allows pushing onto the channel

	c <- 42 //push value onto the channel buffer
	c <- 43 //we can now push this into the channel because we have a buffer of 2. Working with buffers this way is not ideal because you can make mistakes and could be resource costs with large buffers

	fmt.Println(<-c) //executes but only pulls the first item off the buffer
	fmt.Println(<-c) //executes but only pulls the second item off the buffer
}
