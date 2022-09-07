package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"

)

type Coors struct {
	x int
	y int
}

type Ship struct {
	coors       Coors
	orientation int
	size        int
	hits        []Coors
}

type Board struct {
	size  int
	ships []Ship
	hits []Coors
}

func generateShip(size int, board Board) Ship {
	rand.Seed(time.Now().UnixNano())
	orientation := rand.Intn(2)
	var x int
	var y int
	if orientation == 0 {
	coordGeneration:
		for {
			// fmt.Println("COORD GENERATION")
			x = rand.Intn(board.size)
			y = rand.Intn(board.size)
			if x+size > board.size {
				goto coordGeneration
			} else {
				maybeShipCoors := getAllShipCoors(Ship{coors: Coors{x: x, y: y}, orientation: orientation, size: size})
				for _, coors := range maybeShipCoors {
					if isThereShip(board, coors, false) {
						continue coordGeneration
					}
				}
				break coordGeneration
			}
		}
	} else {
	coordGeneration2:
		for {
			x = rand.Intn(board.size)
			y = rand.Intn(board.size)
			if y+size > board.size {
				continue
			} else {
				maybeShipCoors := getAllShipCoors(Ship{coors: Coors{x: x, y: y}, orientation: orientation, size: size})
				for _, coors := range maybeShipCoors {
					if isThereShip(board, coors, false) {
						continue coordGeneration2
					}
				}
				break
			}
		}
	}
	ship := Ship{coors: Coors{x: x, y: y}, orientation: orientation, size: size, hits: []Coors{}}
	return ship
}

func addShips(width2 int, width3 int, width4 int, width5 int, board Board) Board {
	for i := 0; i < width2; i++ {
		ship := generateShip(2, board)
		board.ships = append(board.ships, ship)
	}
	for i := 0; i < width3; i++ {
		ship := generateShip(3, board)
		board.ships = append(board.ships, ship)
	}
	for i := 0; i < width4; i++ {
		ship := generateShip(4, board)
		board.ships = append(board.ships, ship)
	}
	for i := 0; i < width5; i++ {
		ship := generateShip(5, board)
		board.ships = append(board.ships, ship)
	}
	return board
}

func GenerateBoard(size int) Board {
	board := addShips(2, 2, 1, 1, Board{size: size})
	return board
}

func getAllShipCoors(ship Ship) []Coors {
	coors := []Coors{}
	if ship.orientation == 0 {
		for i := 0; i < ship.size; i++ {
			coors = append(coors, Coors{x: ship.coors.x + i, y: ship.coors.y})
		}
	} else {
		for i := 0; i < ship.size; i++ {
			coors = append(coors, Coors{x: ship.coors.x, y: ship.coors.y + i})
		}
	}
	return coors
}

func isThereShip(board Board, coors Coors, debug bool) bool {
	var allOtherShipCoors []Coors
	for _, otherShip := range board.ships {
		allOtherShipCoors = getAllShipCoors(otherShip)
		if debug {
			fmt.Println(allOtherShipCoors)
			fmt.Println(coors)
		}
		for _, shipCoors := range allOtherShipCoors {
			if shipCoors.x == coors.x && shipCoors.y == coors.y {
				if debug {
					fmt.Println("RETURN TRUE")
				}
				return true
			}
		}
	}
	if debug {
		fmt.Println("return false")
	}
	return false
}

func hitBoard(board *Board, coors Coors) Board {
	for i, ship := range board.ships {
		for _, shipCoors := range getAllShipCoors(ship) {
			if shipCoors.x == coors.x && shipCoors.y == coors.y {
				board.ships[i].hits = append(board.ships[i].hits, shipCoors)
				if len(board.ships[i].hits) == board.ships[i].size {
					fmt.Println("SUNK")
				}
			}
		}
	}
	board.hits = append(board.hits, coors)
	return *board
}

func remainingships(board Board) int {
	remaining := 0
	for _, ship := range board.ships {
		if len(ship.hits) != ship.size {
			remaining++
		}
	}
	return remaining
}

func isHit(board Board, coors Coors) bool {
	for _, ship := range board.ships {
		for _, shipHitCoors := range ship.hits {
			if shipHitCoors.x == coors.x && shipHitCoors.y == coors.y {
				return true
			}
		}
	}
	return false
}

func isBoardHit (board Board, coors Coors) bool {
	for _, hitCoors := range board.hits {
		if hitCoors.x == coors.x && hitCoors.y == coors.y {
			return true
		}
	}
	return false
}

func visualizeBoard(board Board) {
	for i := 0; i < board.size; i++ {
		for j := 0; j < board.size; j++ {
			if isThereShip(board, Coors{x: j, y: i}, false) {
				fmt.Printf("⚪|")
			} else {
				fmt.Printf("⚫|")
			}
		}
		fmt.Println()
	}
}




func visualizeHiddenBoard(board Board) string{
	var boardString string
	for i := 0; i < board.size; i++ {
		for j := 0; j < board.size; j++ {
			if isBoardHit(board, Coors{x: j, y: i}) {
				if isHit(board, Coors{x: j, y: i})  {
					boardString += "🟢|"
				} else {
					boardString += "⚫|"
				}
			} else {
				boardString += "⚪️|"
			}
		}
		boardString += "\n"
	}

	//  if ship is sunk 🟠

	// for _, ship := range board.ships {
	// 	if len(ship.hits) == ship.size {
	// 		for _, coors := range getAllShipCoors(ship) {
	// 				if ship.orientation == 0 {
	// 					boardString[coors.x * 2 + coors.y * board.size * 2 + coors.y] = 'H'
	// 				} else {
	// 					boardString[coors.x * 2 + coors.y * board.size * 2 + coors.x] = 'H'
	// 				}
	// 		}
			
	// 	}
	// }

	return boardString
}


func visualizeHiddenBoardHTML(board Board) string{
	var boardString string
	// add header to board string
	boardString += "<!DOCTYPE html> <html> <head> <style> .box { cursor:pointer; } th, td { padding: 5px; text-align: center; } </style> </head> <body>"
	for i := 0; i < board.size; i++ {
		boardString += "<div>"
		for j := 0; j < board.size; j++ {
			if isBoardHit(board, Coors{x: j, y: i}) {
				if isHit(board, Coors{x: j, y: i})  {
					boardString += "<span class=\"box\" onclick=\"hit("+strconv.Itoa(j)+","+strconv.Itoa(i)+")\">🟢</span>"
				} else {
					boardString += "<span class=\"box\" onclick=\"hit("+strconv.Itoa(j)+","+strconv.Itoa(i)+")\">⚫</span>"
				}
			} else {
				boardString += "<span class=\"box\" onclick=\"hit("+strconv.Itoa(j)+","+strconv.Itoa(i)+")\">⚪️</span>"
			}
		}
		boardString += "</div>"
	}
	// send post request to  http://localhost:3001/board/hit
	boardString += "<script> function hit(x, y) { console.log(\"hit\" + x + ' ' + y); var xhttp = new XMLHttpRequest(); xhttp.open(\"POST\", \"http://localhost:3001/hit\", true); xhttp.setRequestHeader(\"Content-type\", \"application/x-www-form-urlencoded\"); xhttp.send(\"x=\"+x+\"&y=\"+y); location.reload();};  </script> </body> </html>"
	return boardString
}
