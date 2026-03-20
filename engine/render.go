package engine

import (
	"fmt"

	"github.com/DevitoDbug/chess_v1/utils"
)

func (e *Engine) RenderBoard() {
	utils.ClearScreen()
	utils.MoveCursorTopLeft()

	// Loop the mf so by row and print the columns
	for row := RowNumber - 1; row >= 0; row-- {
		fmt.Printf("%v %v", utils.Green, row+1)
		fmt.Printf("%v ", utils.White)
		for col := range ColumnNumber {
			piece := e.board[row][col]
			if piece == nil {
				fmt.Printf(" . ")
			} else {
				if piece.Type == King && e.isSquareAttacked(int32(col), int32(row), toggleCurrentPlayer(piece.Color)) {
					fmt.Printf("%v %c %v", utils.Red, GetRenderLetter(piece.Type, piece.Color), utils.White)
				} else {
					fmt.Printf(" %c ", GetRenderLetter(piece.Type, piece.Color))
				}
			}
		}
		fmt.Printf("%v %v", utils.Green, row+1)
		fmt.Printf("%v ", utils.White)
		fmt.Println()
	}

	fmt.Printf("   ")
	for num := range ColumnNumber {
		fmt.Printf("%v %c ", utils.Green, 'a'+num)
	}
	fmt.Printf("%v\n", utils.White)
}

func (e *Engine) RenderEndgame(endGameState EndGameState) {
	e.RenderBoard()

	// TODO: Maybe we find something more interesting to render on an end of a game for now printing text
	// it perfectly fine
	fmt.Println()
	fmt.Println()
	if endGameState.CheckMate {
		fmt.Println("***** Checkmate *****")
		fmt.Printf("%v wins", e.currentPlayerColor)
	} else if endGameState.StaleMate {
		fmt.Println("***** Stalemate *****")
		fmt.Println("It's a draw. GG")
	}
}
