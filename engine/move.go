package engine

import (
	"fmt"

	"github.com/DevitoDbug/chess_v1/utils"
)

func (e *Engine) MovePiece(input Input) error {
	// We need to know the piece that the user has just submitted
	// 	- We can support a switch statement that switches between the pieces, each piece .
	//	- For evaluation movements should include the color and the position that we want to
	// 		move
	switch input.Type {
	case Pawn:
	case King:
	case Queen:
	case Rook:
	case Bishop:
		err := e.MoveBishop(input.StartX, input.StartY, input.DestinationX, input.DestinationY, input.Piece)
		if err != nil {
			return err
		}
	case Knight:
	default:
		return fmt.Errorf("invalid piece provided")
	}
	// 	- We should first find the given piece that wants to be moved.
	//			-> Create a function that will take string given and return the specific piece/array position(row,column)
	// 	- We should have a list of the possible moves for the given piece.
	//	- Things to take note is that.
	// 		-> Position we are moving to could already have something
	//		-> If that things is a piece with the same color as us then we cannot move there
	// 		-> If that thing is of a different color then yes we can move there and fucking replace that give
	// 			piece
	// We need to know what position they intent to move to

	return nil
}

type Move [2]int // Format is (x,y)

func (e *Engine) MovePawn() error {
	possibleMoves := []Move{
		{0, 1}, // Normal move for paws

		// Weird moves (FAAAAAAaaahhhhh)
		{0, 2},  // First move, pawn moves two steps
		{1, 1},  // Ampersand to the right
		{-1, 1}, // Ampersand to the left
	}
	_ = possibleMoves
	return nil
}

func (e *Engine) MoveKing() error {
	possibleMoves := []Move{
		{0, 1},
		{0, -1},
		{-1, 0},
		{1, 0},

		// Diagonals
		{1, 1},
		{-1, 1},
		{-1, -1},
		{1, -1},
	}
	_ = possibleMoves
	return nil
}

func (e *Engine) MoveKnight() error {
	possibleMoves := []Move{
		{-1, 2},
		{1, 2},

		{2, 1},
		{2, -1},

		{-2, 1},
		{-2, -1},

		{-1, -2},
		{1, -2},
	}
	_ = possibleMoves
	return nil
}

func (e *Engine) MoveBishop(startingX, startingY, destinationX, destinationY int, piece *Piece) error {
	// Absolute differences should be same i.e
	// |anotherPointX - pointX | == |anotherPointY - pointY|
	possibleMoves := []Move{}

	// We can take all possible moves for the bishop
	for row := range 8 {
		for col := range 8 {
			// Evaluate magnitude confirms diagonal movement that is |x1 - x2| == |y1 - y2|
			var x1, x2, y1, y2 int
			x1 = row
			y1 = col
			x2 = startingX
			y2 = startingY
			xDiff := utils.AbsoluteDiff(x1, x2)
			yDiff := utils.AbsoluteDiff(y1, y2)

			if x1 == x2 && y1 == y2 {
				// The square we are on is the starting square
				continue
			}

			if xDiff != yDiff {
				continue
			}

			// Check if there is no piece in that square
			square := e.Board[row][col]
			if square == nil {
				possibleMoves = append(possibleMoves, Move{row, col})
				continue
			}

			// Check if the piece in the square is opponents of the current players
			if square.Color == e.CurrentPlayerColor {
				continue
			}

			possibleMoves = append(possibleMoves, Move{row, col})
		}
	}

	for _, possibleMove := range possibleMoves {
		if possibleMove[0] == destinationX && possibleMove[1] == destinationY {
			// Destination moves provided are possible
			e.Board[destinationX][destinationY] = piece

			break
		}
	}

	return fmt.Errorf("move not allowed")
}

func (e *Engine) MoveRook() error {
	return nil
}

func (e *Engine) MoveQueen() error {
	return nil
}
