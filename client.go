package main

import (
	"bufio"
	"fmt"
	"os"
)

type player struct {
	username string
	adress   string
}

func (player *player) initPlayer(username string, adress string) {
	player.username = username
	player.adress = adress
}

func addPlayer() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Hello welcome to the battleship game !\n")

	fmt.Print("Enter a username: \n ")
	username, _ := reader.ReadString('\n')
	fmt.Printf("Player : %s !\n", username)

	fmt.Print("Enter an opponent player: \n ")
	adress, _ := reader.ReadString('\n')
	fmt.Printf("Adress : %s !\n", adress)

	var player player
	player.initPlayer(username, adress)

	//gestion d'erreur : if adress =! adresse Localhost => fmt.Print("Veuillez retaper l'adresse s'il vous plait: \n ")

}

func start() {
	//generateShip
	//addShip
	//generateShip

	//IF NB DE SHIP DU JOEUR EST DIFFERENT DE 0 => if isThereShip(...) = TRUE
	// {

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter a command to attack {username, x, y}: \n")
	attack, _ := reader.ReadString('\n')
	fmt.Printf("Coord : %s !\n", attack)

	//}

	/*
		ELSE {
				fmt.Printf("The battleship is over. Player ... wins the battleship", gagnant.username)
			}
	*/
}

func choice() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter 'start' to start the game, or 'add' to add other player !\n")
	choice, _ := reader.ReadString('\n')

	switch choice {
	case "add":
		addPlayer()
	case "start":
		start()

	default:
		// gestion d'erreur : fmt.Printf("Enter 'start' to start the game, or 'add' to add other player !\n")
	}
}

func play(board Board, myChannel1 chan bool) {
	res := <-myChannel1
	if res {
		addPlayer()
		choice()
	}
}
