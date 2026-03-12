package engine

import (
	"fmt"

	"github.com/DevitoDbug/chess_v1/utils"
)

func (e *Engine) isStraightValidSlidingMove(startingX, startingY, destinationX, destinationY int32) error {
	x1 := startingX
	x2 := destinationX
	y1 := startingY
	y2 := destinationY
	xDiff := x2 - x1
	yDiff := y2 - y1

	// Destination should not have a piece similar to current player color
	if e.Board[destinationY][destinationX] != nil && e.Board[destinationY][destinationX].Color == e.CurrentPlayerColor {
		return fmt.Errorf("move not allowed, piece can not attack a square that has current player's piece")
	}

	xAb := utils.AbsoluteDiff(x1, x2)
	yAb := utils.AbsoluteDiff(y1, y2)
	if xAb == yAb {
		// This eliminates movement to square the piece is already in.
		// Also eliminates diagonal movements.
		return fmt.Errorf("move not allowed, piece can only move in straight lines horizontally or vertically")
	}

	if xAb > 0 && yAb > 0 {
		// This should not be possible, a change in x and y at the same time
		// means the piece did not travel in the expected moves
		return fmt.Errorf("move not allowed, piece can only move in straight lines horizontally or vertically")
	}

	// Loop from the start to the destination and find out if there is another piece in the middle.
	// Below are the only two possibilities, either horizontal or vertical movements.
	if xAb > yAb {
		//-1 is destination pieces should not be detected as obstacles
		for colOffset := range xAb - 1 {
			if xDiff > 0 {
				// Moving to wards the right x increases
				// The +1 is to make sure that we do not check the starting square
				if e.Board[startingY][(startingX+1)+colOffset] != nil {
					return fmt.Errorf("move not allowed, piece path has another piece blocking i.e (%v,%v)", startingX+1+colOffset, startingY)
				}
			} else {
				// Moving to wards the left x decreases
				// The -1 is to make sure that we do not check the starting square
				if e.Board[startingY][(startingX-1)-colOffset] != nil {
					return fmt.Errorf("move not allowed, piece path has another piece blocking i.e (%v,%v)", startingX+1+colOffset, startingY)
				}
			}
		}
	} else if yAb > xAb {
		for rowOffset := range yAb - 1 {
			if yDiff > 0 {
				if e.Board[(startingY+1)+rowOffset][startingX] != nil {
					return fmt.Errorf("move not allowed, piece path has another piece blocking i.e (%v,%v)", startingX, startingY+1+rowOffset)
				}
			} else {
				if e.Board[(startingY-1)-rowOffset][startingX] != nil {
					return fmt.Errorf("move not allowed, piece path has another piece blocking i.e (%v,%v)", startingX, startingY+1+rowOffset)
				}
			}
		}
	}
	return nil
}

func (e *Engine) isDiagonalValidSlidingMove(startingX, startingY, destinationX, destinationY int32) error {
	// Cannot move to a square that has a color similar to the current player
	piece := e.Board[destinationY][destinationX]
	if piece != nil && piece.Color == e.CurrentPlayerColor {
		return fmt.Errorf("move not allowed, destination square has current player's piece")
	}

	x1 := startingX
	x2 := destinationX
	y1 := startingY
	y2 := destinationY

	xAb := utils.AbsoluteDiff(x1, x2)
	yAb := utils.AbsoluteDiff(y1, y2)

	// Diagonal moves should conform to |x1-x2| == |y1-y2|
	if xAb != yAb {
		// Not a valid bishop move
		return fmt.Errorf("move not allowed, move provided is not a diagonal movement")
	}

	// Evaluated diagonal - Basically figuring out if there is another piece along the way
	xChange := x1 - x2
	yChange := y1 - y2
	diagonalX := startingX
	diagonalY := startingY

	// The -1 is, destination pieces should not be detected as obstacles
	for range xAb - 1 {
		// Confirm that nothing is within the diagonal
		if xChange < 0 {
			diagonalX++
		} else {
			diagonalX--
		}

		if yChange < 0 {
			diagonalY++
		} else {
			diagonalY--
		}

		if e.Board[diagonalY][diagonalX] != nil {
			return fmt.Errorf("move not allowed. There is another piece blocking the diagonal i.e at index (%v,%v)", diagonalX, diagonalY)
		}
	}

	return nil
}
