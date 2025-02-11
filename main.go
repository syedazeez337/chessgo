package main

import (
	"fmt"

	"github.com/syedazeez337/chessgo/game"
)

func main() {
	fmt.Println("Welcome to Chess!")

	// Initialize the game
	board := game.NewBoard()
	
	game.PrintBoard(board)
	/*

	// main game loop
	for {
		fmt.Print("Enter your move (e.g., e2 e4):")
		var from, to string
		fmt.Scan(&from, &to)

		// parse and validate the move
		if IsValidMove(board, from, to) {
			MakeMove(board, from, to)
			PrintBoard(board)
		} else {
			fmt.Println("Invalid move. Try again.")
		}
	}
		*/
}
