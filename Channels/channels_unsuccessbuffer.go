package main

import (
	"fmt"
)

func main() {

	c := make(chan int, 1) //here we add 1 to the buffer, this allows us to add "1" value on the channel without needing a reciever at the same time.

	c <- 42 //push value onto the channel buffer
	c <- 43 //here we push a value onto the buffer but is unsuccesful, blocked as the buffer of 1 is already full

	fmt.Println(<-c) //fails to execute do to buffer being full
}
