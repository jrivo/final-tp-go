package main

import (
	"fmt"
	"log"
	"main/battleship/server"
	"net/http"
	"strconv"
	"runtime"
	"os/exec"
	"flag"
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


func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}

func main() {

	port := flag.String("port", "","" )
	flag.Parse()

	board := GenerateBoard(10)
	fmt.Println("Starting web server at 0.0.0.0:"+*port+"...")
	go http.HandleFunc("/board", boardHandler(&board))
	go http.HandleFunc("/boats", boatsHandler(board))
	// send board reference in hit function
	go http.HandleFunc("/hit", hit(&board,1,5))
	myChannel1 := make(chan bool)
	go play(board, myChannel1)
	// openbrowser("http://localhost:3001/board")
	server.Init_server(myChannel1, *port)
}
