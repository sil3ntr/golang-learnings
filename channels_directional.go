package main

import "fmt"

func main() {
	c := make(chan int, 1)    // bi-directional channel
	cs := make(chan<- int, 1) //send channel
	cr := make(<-chan int, 1) //recieve channel

	c <- 42
	cs <- 43
	// cr <- 44 //This would fail because its trying to send on a recieve channel

	fmt.Println("-----")
	fmt.Printf("Is type: %T\n", c)
	fmt.Printf("Is type: %T\n", cs)
	fmt.Printf("Is type: %T\n", cr)
}
