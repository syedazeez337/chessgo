package game

import "fmt"

// type -> pawn, rook; colour -> white, black
type Piece struct {
	Type string
	Colour string
}

type Board [8][8]*Piece

func NewBoard() Board {
	var board Board

	// place white pieces
	board[0] = [8]*Piece {
		{"Rook", "White"},
		{"Knight", "White"},
		{"Bishop", "White"},
		{"Queen", "White"},
		{"King", "White"},
		{"Bishop", "White"},
		{"Knight", "White"},
		{"Rook", "White"},
	}
	for i := 0; i < 8; i++ {
		board[1][i] = &Piece{"Pawn", "White"}
	}

	// place black pieces
	board[7] = [8]*Piece {
		{"Rook", "Black"},
		{"Knight", "Black"},
		{"Bishop", "Black"},
		{"Queen", "Black"},
		{"King", "Black"},
		{"Bishop", "Black"},
		{"Knight", "Black"},
		{"Rook", "Black"},
	}
	for i := 0; i < 8; i++ {
		board[6][i] = &Piece{"Pawn", "Black"}
	}

	return board
}

// Print the board to the console
func PrintBoard(board Board) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] != nil {
				fmt.Printf("%s%s ", board[i][j].Colour[:1], board[i][j].Type[:1])
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}