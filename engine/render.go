package engine

import (
	"fmt"

	"github.com/DevitoDbug/chess_v1/utils"
)

func (e *Engine) RenderTerminal() {
	utils.ClearScreen()
	utils.MoveCursorTopLeft()

	// Loop the mf so by row and print the columns
	for row := range RowNumber {
		fmt.Printf("%v %v", utils.Red, 8-row)
		fmt.Printf("%v |", utils.White)
		for col := range ColumnNumber {
			if e.Board[row][col] == nil {
				fmt.Printf("    |")
			} else {
				fmt.Printf(" %v%v |", GetColorLetter(e.Board[row][col].Color), e.Board[row][col].RenderLetter)
			}
		}
		fmt.Println("")
		for range ColumnNumber {
			fmt.Printf("------")
		}
		fmt.Println("")
	}

	fmt.Printf("   |")
	columnLetters := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for _, letter := range columnLetters {
		fmt.Printf("%v %v  |", utils.Red, letter)
	}
	fmt.Printf("%v\n", utils.White)
}
