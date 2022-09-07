package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var board Board

func boardHandler(w http.ResponseWriter, req *http.Request) {
}

func boatHandler(w http.ResponseWriter, req *http.Request) {
}

func hit(x, y int) func(w http.ResponseWriter, req *http.Request) {
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
			board = hitBoard(board,coors);
			fmt.Fprintf(w, "you hit the board at %d,%d", x, y);
	}
}
return hitHandler
}



func main() {
	board = GenerateBoard(10)
	fmt.Println("Starting web server at 0.0.0.0:3001...")
	http.HandleFunc("/board", boardHandler)
	http.HandleFunc("/boats", boatHandler)
	http.HandleFunc("/hit", hit(1,5))

	if err := http.ListenAndServe("0.0.0.0:3001", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
