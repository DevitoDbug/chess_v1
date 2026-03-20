package engine

import (
	"fmt"
)

// Run - has the game loop
func (e *Engine) Run() {
	e.RenderBoard()

	for {
		var input string

		fmt.Printf("Enter your move(%v)\n: ", e.currentPlayerColor)
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Invalid value entered")
		}

		parsedInput, err := ParseInput(input)
		if err != nil {
			fmt.Println("Invalid value entered")
		}
		fmt.Printf("parse Input is: %+v\n", parsedInput)

		err = e.MovePiece(parsedInput)
		if err != nil {
			fmt.Printf("%v\n", err)
			e.RenderBoard()
			continue
		}

		// Check for checkmate or stalemate
		endGameState := e.GetEndGameState()
		if endGameState != nil {
			e.RenderEndgame(*endGameState)
			break
		}

		e.currentPlayerColor = toggleCurrentPlayer(e.currentPlayerColor)

		fmt.Println("******************************************")
		fmt.Println("******************************************")

		e.RenderBoard()
	}
}
