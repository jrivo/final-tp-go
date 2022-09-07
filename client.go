package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type OpponentPlayer struct {
	username string
	adress   string
}

type mainPlayer struct {
	username  string
	opponents []OpponentPlayer
}

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
			if scanner.Text() == "" {
				fmt.Println("Invalid username")
				continue opponentCreationLoop
			}
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

func parseCommand(command string) (string, int, int) {
	var username string
	var targetX int
	var targetY int
	elements := strings.Split(command, " ")
	username = elements[0]
	//TODO complete command to handle just username
	targetX, _ = strconv.Atoi(elements[1])
	targetY, _ = strconv.Atoi(elements[2])
	return username, targetX, targetY
}

func attackOpponent(opponent OpponentPlayer, targetX int, targetY int) {
	_, err := http.NewRequest(http.MethodPost, "http://"+opponent.adress+"/hit", nil)
	if err == nil {
		fmt.Println("Attack successfull")
	} else {
		fmt.Println("Error", err)
	}
}

func startGame(player mainPlayer) {
	//generateShip
	//addShip
	//generateShip

	//IF NB DE SHIP DU JOEUR EST DIFFERENT DE 0 => if isThereShip(...) = TRUE
	// {
	// fmt.Println(player.opponents)
	scanner := bufio.NewScanner(os.Stdin)
commandLoop:
	for {
		var targetUsername string
		var targetPlayer OpponentPlayer
		var targetX int
		var targetY int
		fmt.Printf("Enter a command to attack {username, x, y}: \n")
		if scanner.Scan() {
			if scanner.Text() != "" {
				if (len(strings.Split(scanner.Text(), " "))) != 3 {
					fmt.Println("Invalid command")
					continue commandLoop
				}
				targetUsername, targetX, targetY = parseCommand(scanner.Text())
				fmt.Printf("Username: %s, X: %d, Y: %d \n", targetUsername, targetX, targetY)
			opponentExistLoop:
				for _, opponent := range player.opponents {
					fmt.Println(opponent.username, targetUsername)
					if opponent.username == targetUsername {
						targetPlayer = opponent
						break opponentExistLoop
					}
					fmt.Println("Opponent doesn't exist")
					continue commandLoop
				}
				attackOpponent(targetPlayer, targetX, targetY)
				fmt.Printf("Attacking %s at %d, %d \n", targetUsername, targetX, targetY)
			} else {
				fmt.Println("Invalid username")
			}
			continue
		}
		// fmt.Printf("Coord : %s !\n", attack)

		//}

		/*
			ELSE {
					fmt.Printf("The battleship is over. Player ... wins the battleship", gagnant.username)
				}
		*/
	}
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
