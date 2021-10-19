package main

import (
	"fmt"
)

func main() {

	c := make(chan int, 1) //here we add 1 to the buffer, this allows us to add "1" value on the channel without needing a reciever at the same time.

	c <- 42 //push value onto the channel buffer

	fmt.Println(<-c) //pull the value off the channel
}
