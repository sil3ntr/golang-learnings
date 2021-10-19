package main

import (
	"fmt"
	"../learn"
)

func main() {

	p := `Password123#`

	bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
	fmt.Println(bs)

}
