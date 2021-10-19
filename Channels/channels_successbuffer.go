package main

import (
	"fmt"
)

func main() {

	c := make(chan int, 2) //here we add 1 to the buffer, this allows us to add "1" value on the channel without needing a reciever at the same time.

	c <- 42 //push value onto the channel buffer
	c <- 43 //we can now push this into the channel because we have a buffer of 2. Working with buffers this way is not ideal because you can make mistakes and could be resource costs with large buffers

	fmt.Println(<-c) //executes but only pulls the first item off the buffer
	fmt.Println(<-c) //executes but only pulls the second item off the buffer
}
