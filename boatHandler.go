package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BoatResponse struct {
	NbrBateau int
}

// GET /boats:
// Devra renvoyer le nombre de bateaux qui sont encore à flot. Si le serveur renvoie 0, cela
// signifie que ce joueur a perdu.

func boatsHandler(board Board) func(w http.ResponseWriter, req *http.Request) {
	boatsHandlerReturn := func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			fmt.Println("GET /boats")
			boatResp := BoatResponse{len(board.ships)}
			w.Header().Set("content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(boatResp)
		}
	}
	return boatsHandlerReturn
}
