package engine

import "fmt"

// This file contains the game loop

func (e *Engine) Run() {
	for {
		fmt.Println("hello")

		// Render the board
		e.RenderTerminal()

		var input string

		fmt.Println("Enter your move: ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Invalid value entered")
		}
		// Pick input from the user
		// Process the input
		err = e.MovePiece(input)
		if err != nil {
			println("error")
		}
		fmt.Println("******************************************")
	}
}
