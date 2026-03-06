package engine

import "fmt"

func (e *Engine) MovePiece(moveCode string) error {
	// We need to know the piece that the user has just submitted
	// 	- We can support a switch statement that switches between the pieces, each piece .
	//	- For evaluation movements should include the color and the position that we want to
	// 		move
	// 	- We should first find the given piece that wants to be moved.
	//			-> Create a function that will take string given and return the specific piece/array position(row,column)
	// 	- We should have a list of the possible moves for the given piece.
	//	- Things to take note is that.
	// 		-> Position we are moving to could already have something
	//		-> If that things is a piece with the same color as us then we cannot move there
	// 		-> If that thing is of a different color then yes we can move there and fucking replace that give
	// 			piece
	// We need to know what position they intent to move to

	fmt.Println(moveCode)
	return nil
}
