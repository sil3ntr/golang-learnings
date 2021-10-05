package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World!")
		d, _ := ioutil.ReadAll(r.Body)
		log.Printf("Data %s\n", d)
	})

	http.ListenAndServe(":9090", nil)
}
