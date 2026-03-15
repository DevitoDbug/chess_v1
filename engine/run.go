package engine

import (
	"fmt"
)

// This file contains the game loop

func (e *Engine) Run() {
	e.RenderTerminal()

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

		// Check if it is a valid move
		// 	-> Is it the correct correct color
		//	-> Is the destination allowed (Not out of bound, our piece can move there, another of our piece is not there)
		err = e.MovePiece(parsedInput)
		if err != nil {
			fmt.Printf("%v\n", err)
		} else {
			e.currentPlayerColor = toggleCurrentPlayer(e.currentPlayerColor)
		}

		fmt.Println("******************************************")
		fmt.Println("******************************************")

		e.RenderTerminal()
	}
}
