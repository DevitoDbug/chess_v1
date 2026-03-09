package engine

import (
	"fmt"
)

// This file contains the game loop

func (e *Engine) Run() {
	for {
		// Render the board
		e.RenderTerminal()

		var input string

		fmt.Println("Enter your move: ")
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
			fmt.Printf("%v", err)
		}

		fmt.Println("******************************************")
		fmt.Println("******************************************")
	}
}
