package main

import (
	"encoding/json"
	"fmt"
)

type computer struct {
	ModelName string `json: "ModelName"`
	ModelNum  string `json: "ModelNum"`
	ManuFact  string `json: "ManuFact"`
	Mem       int    `json: "Mem"`
	Rrp       int    `json: "Rrp"`
}

func main() {

	jsonBlob := []byte(`[
		{"ModelName": "Amiga 500", "ModelNum": "500", "ManuFact": "Commadore", "Mem": 500, "Rrp": 699},
		{"ModelName": "Spectrum 128K+", "ModelNum": "128k", "ManuFact": "Sinclair", "Mem": 128, "Rrp": 499}
	]`)

	computers := []computer{}
	// or var computers []computer

	err := json.Unmarshal(jsonBlob, &computers)

	if err != nil {
		fmt.Println("Error is:", err)
	}

	fmt.Printf("This is the value: %v \n", computers)
	fmt.Printf("This is the Type: %T \n", computers)

	for i, v := range computers {
		fmt.Printf("Computer Index is # %v \n", i)
		fmt.Printf("\tThe Computer Model is: %v \n", v.ModelName)
		fmt.Printf("\tIt has %v kb of memory \n", v.Mem)
		fmt.Printf("\tIt Cost $ %v RRP when first Launched\n", v.Rrp)
		fmt.Printf("\tIt was made by %v \n", v.ManuFact)
	}

}
