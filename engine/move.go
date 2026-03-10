package engine

import (
	"fmt"

	"github.com/DevitoDbug/chess_v1/utils"
)

func (e *Engine) MovePiece(input Input) error {
	piece := e.Board[input.StartY][input.StartX]
	if piece == nil {
		return fmt.Errorf("no piece in the square referenced")
	}

	// Catch players moving other opponents pieces early
	if piece.Color != e.CurrentPlayerColor {
		return fmt.Errorf("invalid move, you can only move pieces belonging to the current player")
	}

	switch piece.Type {
	case Pawn:
		err := e.MovePawn(input.StartX, input.StartY, input.DestinationX, input.DestinationY)
		if err != nil {
			return err
		}
	case King:
		return fmt.Errorf("king move not implemented yet")
	case Queen:
		return fmt.Errorf("queen move not implemented yet")
	case Rook:
		return fmt.Errorf("rook move not implemented yet")
	case Bishop:
		err := e.MoveBishop(input.StartX, input.StartY, input.DestinationX, input.DestinationY)
		if err != nil {
			return err
		}
	case Knight:
		return fmt.Errorf("night move not implemented yet")
	default:
		return fmt.Errorf("invalid piece provided")
	}

	// Clean up AmpersandSquare
	if e.AmpersandSquare != nil {
		if e.AmpersandSquare.Checked {
			e.AmpersandSquare = nil
		} else {
			e.AmpersandSquare.Checked = true
		}
	}

	return nil
}

type Move [2]int // Format is (x,y)

func (e *Engine) MovePawn(startingX, startingY, destinationX, destinationY int32) error {
	startingPiece := e.Board[startingY][startingX]
	destinationPiece := e.Board[destinationY][destinationX]

	x1 := startingX
	y1 := startingY
	x2 := destinationX
	y2 := destinationY
	xAbsDiff := utils.AbsoluteDiff(x1, x2)
	yAbsDiff := utils.AbsoluteDiff(y1, y2)

	xDiff := x2 - x1
	yDiff := int32(0)
	switch e.CurrentPlayerColor {
	// White moves are positive while black moves are negative
	// White moves from a 0 -> 7 while black moves from 7 -> 0
	case White:
		yDiff = y2 - y1
	case Black:
		yDiff = y1 - y2
	}

	// Attacking a piece
	if xAbsDiff == 1 && yAbsDiff == 1 && yDiff > 0 {
		// yDiff check is to make sure that we are moving in the right direction,
		// white is not moving toward 0 and black towards 7
		isAmpersandMove := e.AmpersandSquare != nil && destinationY == e.AmpersandSquare.RowIndex && destinationX == e.AmpersandSquare.ColumnIndex
		isAttackingOpponentPiece := destinationPiece != nil && destinationPiece.Color != e.CurrentPlayerColor
		if !isAttackingOpponentPiece && !isAmpersandMove {
			return fmt.Errorf("player is not allowed to attack his own piece")
		}

		if isAmpersandMove {
			switch e.CurrentPlayerColor {
			case White:
				e.Board[e.AmpersandSquare.RowIndex-1][e.AmpersandSquare.ColumnIndex] = nil
			case Black:
				e.Board[e.AmpersandSquare.RowIndex+1][e.AmpersandSquare.ColumnIndex] = nil
			}
		}

		e.Board[destinationY][destinationX] = e.Board[startingY][startingX]
		e.Board[startingY][startingX] = nil
		return nil
	}

	if xDiff != 0 {
		return fmt.Errorf("pawns are only allowed to move forward or attack diagonally one step")
	}

	if yDiff == 0 {
		return fmt.Errorf("pawn did not move")
	}

	if yDiff != 2 && yDiff != 1 {
		return fmt.Errorf("pawns are only allowed to move one step forward, or two steps for the first move")
	}

	if yDiff == 2 {
		switch startingPiece.Color {
		// Moving two steps creates possibility for ampersand in the position (starting position + 1) or
		// (destination position -1 ) from whites perspective
		case White:
			if startingY != 1 {
				return fmt.Errorf("paws can only move two moves if it is the first move")
			}
			e.AmpersandSquare = &Square{
				RowIndex:    int32(startingY + 1),
				ColumnIndex: startingX,
			}
		case Black:
			if startingY != 6 {
				return fmt.Errorf("paws can only move two moves if it is the first move")
			}
			e.AmpersandSquare = &Square{
				RowIndex:    int32(startingY - 1),
				ColumnIndex: startingX,
			}
		default:
			return fmt.Errorf("pawn color not recognized")
		}
	}

	if e.Board[destinationY][destinationX] != nil {
		return fmt.Errorf("pawns cannot move to squares that have pieces")
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
	piece := e.Board[destinationY][destinationX]
	if piece != nil && piece.Color == e.CurrentPlayerColor {
		return fmt.Errorf("move not allowed, destination square has current player's piece")
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
