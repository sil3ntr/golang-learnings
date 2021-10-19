package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 42 //send 42 onto chan c
		close(c)
	}()
	v, ok := <-c //pulls value off chan c
	fmt.Println(v, ok)

	v, ok = <-c //trys to pull value off chan c, fails
	fmt.Println(v, ok)
}
