package engine

import (
	"fmt"

	"github.com/DevitoDbug/chess_v1/utils"
)

func (e *Engine) RenderTerminal() {
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
