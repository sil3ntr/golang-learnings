package main

import (
	"fmt"
	"strconv"
)

func main() {
	bitSize := 32
	fmt.Printf(" %v, %T\n", bitSize, bitSize)
	//var j string
	j := strconv.Itoa(bitSize)
	fmt.Printf(" %v, %T\n", j, j)
}
