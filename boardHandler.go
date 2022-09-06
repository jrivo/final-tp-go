package main

import (
	"fmt"
	"net/http"
)

func BoardHandler(board Board) func(w http.ResponseWriter, req *http.Request) {
	boardHandlerReturn := func(w http.ResponseWriter, req *http.Request) {
		fmt.Println(board)
		switch req.Method {
		case http.MethodGet:
			fmt.Fprintf(w, "GOOD\n")
			fmt.Fprintf(w, "Size %d", board.size)
		}
	}
	return boardHandlerReturn
}
