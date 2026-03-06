package engine

import (
	"fmt"
)

func (e *Engine) RenderTerminal() {
	// Loop the mf so by row and print the columns
	for row := range RowNumber {
		for col := range ColumnNumber {
			if e.Board[row][col] == nil {
				fmt.Printf("    |")
			} else {
				fmt.Printf(" %v%v |", GetColorLetter(e.Board[row][col].Color), e.Board[row][col].RenderLetter)
			}
		}
		fmt.Println("")
		for range ColumnNumber {
			fmt.Printf("-----")
		}
		fmt.Println("")
	}
}
