package main

import (
	"fmt"
)

func main() {

	c := make(chan int) // general bidirectional channel expecting int
	go foo(c)           //send - call foo and provide c channel

	go bar(c) //receive -

}

func foo(c chan<- int) { //Local scope of channel here, possible because the general channel type created and passed in. This is now a send only channel.
	c <- 42
	fmt.Printf("%T \n", c)

} //send

func bar(c <-chan int) { //This is now a recieve only channel
	fmt.Println(<-c)
} //recieve
