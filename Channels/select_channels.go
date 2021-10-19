package main

func main() {

	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	//send

	go send(eve, odd, quit)

	//recieve
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
	q <- 0
}
