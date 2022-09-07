package main

import (
	"bufio"
	"fmt"
	"os"
)

type OpponentPlayer struct {
	username string
	adress   string
}

type mainPlayer struct {
	username  string
	opponents []OpponentPlayer
}

// type Opponents struct {
// 	players []Player
// }

// func (player *Player) initPlayer(username string, adress string) {
// 	player.username = username
// 	player.adress = adress
// }

func addOpponent(player *mainPlayer) {
	var opponentPlayer OpponentPlayer
	var opponentUsername string
	var opponentAdress string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter an opponent player: \n ")
opponentCreationLoop:
	for {
		fmt.Print("Enter a username: \n ")
		if scanner.Scan() {
			opponentUsername = scanner.Text()
			for {
				fmt.Print("Enter an adress: \n ")
				if scanner.Scan() {
					opponentAdress = scanner.Text()
					break opponentCreationLoop
				}
				fmt.Println("Invalid adress")
			}
		}
		fmt.Println("Invalid username")
	}
	opponentPlayer = OpponentPlayer{username: opponentUsername, adress: opponentAdress}
	player.opponents = append(player.opponents, opponentPlayer)
	//gestion d'erreur : if adress =! adresse Localhost => fmt.Print("Veuillez retaper l'adresse s'il vous plait: \n ")

}

func startGame(player mainPlayer) {
	//generateShip
	//addShip
	//generateShip

	//IF NB DE SHIP DU JOEUR EST DIFFERENT DE 0 => if isThereShip(...) = TRUE
	// {
	fmt.Println(player.opponents)
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

func play(board Board) {
	var username string
	var player mainPlayer
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Hello welcome to the battleship game !\n")
	for {
		fmt.Print("Enter a username: \n ")
		if scanner.Scan() {
			if scanner.Text() != "" {
				username = scanner.Text()
				player = mainPlayer{username: username}
				break
			}
			fmt.Println("Invalid username")
			continue
		}
	}
	fmt.Printf("Player : %s !\n", username)
addOrStartLoop:
	for {
		addOpponent(&player)
		fmt.Println("Enter 'start' to start the game or 'add' to add a player")
		if scanner.Scan() {
			switch scanner.Text() {
			case "start":
				break addOrStartLoop
			case "add":
				continue addOrStartLoop
			default:
				fmt.Println("Invalid command")
				continue addOrStartLoop
			}
		}

	}
	startGame(player)
}
