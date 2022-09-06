package main

import (
	"bufio"
	"fmt"
	"os"
)

type player struct {
	username []string
	adress   []string
	board    Board
}

func addPlayer() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Hello welcome to the battleship game !\n")

	fmt.Print("Enter a username ")
	username, _ := reader.ReadString('\n')
	fmt.Printf("Player : %s !\n", username)

	fmt.Print("Enter an opponent player: ")
	adress, _ := reader.ReadString('\n')
	fmt.Printf("Adress : %s !\n", adress)
}

func start() {
	//generateShip
	//addShip
	//generateShip

	//IF NB DE SHIP DU JOEUR EST DIFFERENT DE 0 => if isThereShip(...) = TRUE
	// {

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter a command to attack {username, x, y}\n")
	attack, _ := reader.ReadString('\n')
	fmt.Printf("Coord : %s !\n", attack)

	//}

	/*
		ELSE {
				fmt.Printf("Le bataille navale est terminé. Le joueur %s à gagné la partie !\n", gagnant.username)
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
	default:
		start()
	}
}
