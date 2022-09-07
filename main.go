package main

import (
	"fmt"
	"log"
	"main/battleship/server"
	"net/http"
	"strconv"
)


func hit(board *Board,x int , y int) func(w http.ResponseWriter, req *http.Request) {
	hitHandler := func(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
		case http.MethodGet:
		fmt.Fprintf(w, "you can attack using a post request with x and y coordinates");
		case http.MethodPost:
			// get coordinates
			var coors Coors
			x, err :=  strconv.Atoi(req.FormValue("x"))
			y, err :=  strconv.Atoi(req.FormValue("y"))
			if err != nil {
				log.Fatal(err)
			}
			coors.x = x
			coors.y = y
			*board = hitBoard(board,coors);
			fmt.Fprintf(w, "you hit the board at %d,%d", x, y);
			// print the board
			fmt.Println(board)
	}
}
return hitHandler
}


func main() {

	board := GenerateBoard(10)
	fmt.Println("Starting web server at 0.0.0.0:3001...")
	go http.HandleFunc("/board", boardHandler(&board))
	go http.HandleFunc("/boats", boatsHandler(board))
	// send board reference in hit function
	go http.HandleFunc("/hit", hit(&board,1,5))
	myChannel1 := make(chan bool)
	go play(board, myChannel1)

	server.Init_server(myChannel1)
}
