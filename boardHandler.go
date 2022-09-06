package main

import (
	"fmt"
	"net/http"
)

// GET /board:
// Devra renvoyer l’état du plateau. Un exemple de format possible peut être plusieurs suites
// de chiffres:
// - 0: case intouchée, non dévoilée
// - 1: coup dans l’eau
// - 2: bateau touché
// Libre à vous de choisir un meilleur format de communication.

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
