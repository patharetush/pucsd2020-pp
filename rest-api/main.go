package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!\n", r.URL.Path[1:])
}

func main() {
	log.Println("Starting Simple Web-Server")
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if nil != err {
		log.Fatal("Failed to start web-server on port 8080: ", err.Error())
	}
}
