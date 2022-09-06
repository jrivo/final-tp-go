package main

import (
	"fmt"
	"log"
	"net/http"
)

func boardHandler(w http.ResponseWriter, req *http.Request) {
}

func boatHandler(w http.ResponseWriter, req *http.Request) {
}

func hitHandler(w http.ResponseWriter, req *http.Request) {}

func main() {
	fmt.Println("Starting web server...")
	http.HandleFunc("/board", boardHandler)
	http.HandleFunc("/boats", boatHandler)
	http.HandleFunc("/hit", hitHandler)

	if err := http.ListenAndServe("0.0.0.0:3001", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
