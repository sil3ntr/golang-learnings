package main

import (
	"fmt"
)

type person struct {
	first string
	age   int
}

func (*person) speak() {
	fmt.Println("This is method speak")
}

type human interface {
	speak()
}

func saySomething(h human) {
	fmt.Printf("This is func saySomething %T \n", h)
	fmt.Printf("This is h in saySomething: %v\n", h)
	h.speak()
}

func main() {

	fmt.Println("Check base is working")

	p1 := person{
		first: "Rich",
		age:   40,
	}
	//fmt.Println(p1)
	saySomething(&p1)

}
