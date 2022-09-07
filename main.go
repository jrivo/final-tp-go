package main

import (
	"fmt"
	"log"
	"main/battleship/server"
	"net/http"
	"strconv"
)

var board Board

func hit(x, y int) func(w http.ResponseWriter, req *http.Request) {
	hitHandler := func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			fmt.Fprintf(w, "you can attack using a post request with x and y coordinates")
		case http.MethodPost:
			// get coordinates
			var coors Coors
			x, err := strconv.Atoi(req.FormValue("x"))
			y, err := strconv.Atoi(req.FormValue("y"))
			if err != nil {
				log.Fatal(err)
			}
			coors.x = x
			coors.y = y
			board = hitBoard(board, coors)
			fmt.Fprintf(w, "you hit the board at %d,%d", x, y)
		}
	}
	return hitHandler
}

func main() {
	fmt.Println("----- Board Game -----")
	board = GenerateBoard(10)

	http.HandleFunc("/board", boardHandler())
	fmt.Println("/board")

	http.HandleFunc("/boats", boatsHandler(board))
	fmt.Println("/boats")

	http.HandleFunc("/hit", hit(1, 5))
	fmt.Println("/hit")

	myChannel1 := make(chan bool)
	go play(board, myChannel1)

	server.Init_server(myChannel1)
}
