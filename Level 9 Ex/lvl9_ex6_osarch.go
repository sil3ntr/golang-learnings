package main

import (
	"fmt"
	"runtime"
)

func main() {
	//fmt.Println("hello")
	fmt.Println("CPU Archiecture:", runtime.GOARCH)
	fmt.Println("Operating System:", runtime.GOOS)
}
