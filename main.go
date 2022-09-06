package main

import (
	"fmt"
	"log"
	"net/http"
)

func boatHandler(w http.ResponseWriter, req *http.Request) {
}

func hitHandler(w http.ResponseWriter, req *http.Request) {}

func main() {
	fmt.Println("Starting web server...")
	board := GenerateBoard(10)
	http.HandleFunc("/board", BoardHandler(board))
	http.HandleFunc("/boats", boatHandler)
	http.HandleFunc("/hit", hitHandler)

	go PlayGame(board)

	if err := http.ListenAndServe("0.0.0.0:3001", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
