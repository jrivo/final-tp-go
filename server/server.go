package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func Init_server(myChannel1 chan bool) {
	godotenv.Load(".env")
	fmt.Println("Starting web server...")
	fmt.Println("Server listen to port", os.Getenv("PORT_SERVER"))
	myChannel1 <- true
	http.ListenAndServe(":"+os.Getenv("PORT_SERVER"), nil)
}
