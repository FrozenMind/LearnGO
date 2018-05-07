package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var playerTurn bool = true //true = player1 (X), false = player2 (O)
var board = [3][3]string{{" ", " ", " "}, {" ", " ", " "}, {" ", " ", " "}}

func main() {

	reader := bufio.NewReader(os.Stdin)

	printField(board)

	var x, y int
	//var input string
	for {
		fmt.Print("Choose your field (row,column): ")
		input, _ := reader.ReadString('\n')
		fields := strings.Split(input, "")
		x = strToInt(fields[0])
		y = strToInt(fields[1])
		if board[x][y] != "X" && board[x][y] != "O" {
			if playerTurn == true {
				board[x][y] = "X"
				playerTurn = false
				fmt.Println("Player 1 places on", x, y)
			} else {
				board[x][y] = "O"
				playerTurn = true
				fmt.Println("Player 2 places on", x, y)
			}
		} else {
			fmt.Println("Field is already used. Pleasy choose another one")
		}
		printField(board)
		if checkWin(x, y) == true {
			//last player, is converted here
			if playerTurn == true {
				fmt.Println("Player 2 win")
			} else {
				fmt.Println("Player 1 win")
			}

			os.Exit(0)
		}
	}
}

func checkWin(x, y int) bool {
	//check vertical line
	if board[0][y] == board[1][y] && board[1][y] == board[2][y] {
		return true
	}
	//check horizontal line
	if board[x][0] == board[x][1] && board[x][1] == board[x][2] {
		return true
	}
	//check if previous move was on the main diagonal and caused a win
	if x == y && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return true
	}
	//check if previous move was on the secondary diagonal and caused a win
	if x+y == 2 && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return true
	}
	//no one has won yet
	return false
}
