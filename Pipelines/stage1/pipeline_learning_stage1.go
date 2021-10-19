package main

import "fmt"

//trying to understand the concept https://go.dev/blog/pipelines

func main() {
	//pipeline setup
	c := gen(2, 3) //ints going into func gen

	out := sq(c) //passing the send channel

	// receiving - consuming the pipeline output

	fmt.Println(<-out) // expected value 4 - ie 3*2 from sq
	fmt.Println(<-out) // expected value 9 - ie 3*3 from sq

}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums { // reads through ints in nums and sends them on the out channel
			out <- n
		}
		close(out) // closing the out send channel
	}()
	return (out) // remember to return the send channel
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n //squaring numbers being sent over close channel from func gen and passing them onto the out send channel
		}
		close(out)
	}()
	return (out)
}
