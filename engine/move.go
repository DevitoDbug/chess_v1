package engine

import (
	"fmt"

	"github.com/DevitoDbug/chess_v1/utils"
)

func (e *Engine) MovePiece(input Input) error {
	// Evaluate the type of piece passed in
	fmt.Println(e)
	fmt.Printf("%+v\n", input)

	piece := e.Board[input.StartY][input.StartX]
	if piece == nil {
		return fmt.Errorf("no piece in the square referenced")
	}

	switch piece.Type {
	case Pawn:
		err := e.MovePawn(input.StartX, input.StartY, input.DestinationX, input.DestinationY)
		if err != nil {
			return err
		}
	case King:
	case Queen:
	case Rook:
	case Bishop:
		err := e.MoveBishop(input.StartX, input.StartY, input.DestinationX, input.DestinationY)
		if err != nil {
			return err
		}
	case Knight:
	default:
		return fmt.Errorf("invalid piece provided")
	}
	return nil
}

type Move [2]int // Format is (x,y)

func (e *Engine) MovePawn(startingX, startingY, destinationX, destinationY int32) error {
	piece := e.Board[startingY][startingX]

	if piece.Color != e.CurrentPlayerColor {
		return fmt.Errorf("invalid move, you can only move pieces belonging to the current player")
	}

	_ = []Move{
		{0, 1}, // Normal move for pawns

		// Weird moves (FAAAAAAaaahhhhh)
		{0, 2},  // First move, pawn moves two steps
		{1, 1},  // Ampersand to the right
		{-1, 1}, // Ampersand to the left
	}
	// TODO: Ampersand

	x1 := startingX
	y1 := startingY
	x2 := destinationX
	y2 := destinationY

	xDiff := x2 - x1
	if xDiff != 0 {
		return fmt.Errorf("currently pawns are only allowed to move forward")
	}

	yDiff := y2 - y1
	if yDiff != 2 && yDiff != 1 {
		return fmt.Errorf("pawns are only allowed to move one step forward, or two steps for the first move")
	}

	if yDiff == 2 {
		switch piece.Color {
		case White:
			if startingY != 1 {
				return fmt.Errorf("paws can only move two moves if it is the first move")
			}
		case Black:
			if startingY != 6 {
				return fmt.Errorf("paws can only move two moves if it is the first move")
			}
		default:
			return fmt.Errorf("pawn color not recognized")
		}
	}

	e.Board[destinationY][destinationX] = e.Board[startingY][startingX]
	e.Board[startingY][startingX] = nil

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

func (e *Engine) MoveBishop(startingX, startingY, destinationX, destinationY int32) error {
	// Cannot move to a square that has a color similar to the current player
	if e.Board[destinationY][destinationX] != nil && e.Board[destinationY][destinationX].Color == e.CurrentPlayerColor {
		return fmt.Errorf("move not allowed")
	}

	x1 := startingX
	x2 := destinationX
	y1 := startingY
	y2 := destinationY

	xAd := utils.AbsoluteDiff(x1, x2)
	yAd := utils.AbsoluteDiff(y1, y2)

	// Diagonal moves should conform to |x1-x2| == |y1-y2|
	if xAd != yAd {
		// Not a valid bishop move
		return fmt.Errorf("move not allowed")
	}

	// Evaluated diagonal - Basically figuring out if there is a piece along the way
	xChange := x1 - x2
	yChange := y1 - y2
	for range xAd {
		diagonalX := startingX
		diagonalY := startingY

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
			return fmt.Errorf("move not allowed. There is a piece blocking the diagonal i.e at index (%v,%v)", diagonalX, diagonalY)
		}
	}

	e.Board[destinationY][destinationX] = e.Board[startingY][startingX]
	e.Board[startingY][startingX] = nil
	return nil
}

func (e *Engine) MoveRook() error {
	return nil
}

func (e *Engine) MoveQueen() error {
	return nil
}
