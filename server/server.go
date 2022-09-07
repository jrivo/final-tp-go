package server

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
)

func Init_server(myChannel1 chan bool, port string) {
	godotenv.Load(".env")
	fmt.Println("Starting web server...")
	fmt.Println("Server listen to port", port)
	myChannel1 <- true
	http.ListenAndServe(":"+port, nil)
}
