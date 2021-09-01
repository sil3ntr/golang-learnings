package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func main() {
	//Read the JSON file
	content, err := ioutil.ReadFile("./wf.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	//Unmarshall the data into the variable `payload`
	var payload map[string]interface{}
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	//Lets print the unmarshalled data
	log.Printf("links: %s\n", payload["links"])
	log.Printf("meta: %s\n", payload["meta"])
	log.Printf("data: %s\n", payload["data"])

}
