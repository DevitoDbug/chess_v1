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
		err := e.MoveRook(input.StartX, input.StartY, input.DestinationX, input.DestinationY)
		if err != nil {
			return err
		}
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

	// Clean ups
	if e.EnpassantSquare != nil {
		if e.EnpassantSquare.Checked {
			e.EnpassantSquare = nil
		} else {
			e.EnpassantSquare.Checked = true
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
		isEnpassantMove := e.EnpassantSquare != nil && destinationY == e.EnpassantSquare.RowIndex && destinationX == e.EnpassantSquare.ColumnIndex
		isAttackingOpponentPiece := destinationPiece != nil && destinationPiece.Color != e.CurrentPlayerColor
		if !isAttackingOpponentPiece && !isEnpassantMove {
			return fmt.Errorf("player is not allowed to attack his own piece")
		}

		if isEnpassantMove {
			switch e.CurrentPlayerColor {
			case White:
				e.Board[e.EnpassantSquare.RowIndex-1][e.EnpassantSquare.ColumnIndex] = nil
			case Black:
				e.Board[e.EnpassantSquare.RowIndex+1][e.EnpassantSquare.ColumnIndex] = nil
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
			e.EnpassantSquare = &Square{
				RowIndex:    int32(startingY + 1),
				ColumnIndex: startingX,
			}
		case Black:
			if startingY != 6 {
				return fmt.Errorf("paws can only move two moves if it is the first move")
			}
			e.EnpassantSquare = &Square{
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
	// The -1 is to make sure that we only check for piece between the staring and the destination and
	// not the specific destination, otherwise it would be detected as a piece in the path.
	for range xAd - 1 {
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

func (e *Engine) MoveRook(startingX, startingY, destinationX, destinationY int32) error {
	x1 := startingX
	x2 := destinationX
	y1 := startingY
	y2 := destinationY
	xDiff := x2 - x1
	yDiff := y2 - y1

	// Destination should not have a piece similar to current player color
	if e.Board[destinationY][destinationX] != nil && e.Board[destinationY][destinationX].Color == e.CurrentPlayerColor {
		return fmt.Errorf("move not allowed, rook can not attack a square that has current player's piece")
	}

	xAd := utils.AbsoluteDiff(x1, x2)
	yAd := utils.AbsoluteDiff(y1, y2)
	if xAd == yAd {
		// This eliminates movement to square the piece is already in.
		// Also eliminates diagonal movements.
		return fmt.Errorf("move not allowed, rook can only move in straight lines horizontally or vertically")
	}

	if xAd > 0 && yAd > 0 {
		// This should not be possible, a change in x and y at the same time
		// means the rook did not travel in the expected moves
		return fmt.Errorf("move not allowed, rook can only move in straight lines horizontally or vertically")
	}

	// Loop from the start to the destination and find out if there is a piece in the middle.
	// Below are the only two possibilities, either horizontal or vertical movements.
	if xAd > yAd {
		//-1 is to make sure we check only up to the destination square and not the destination square itself
		for colOffset := range xAd - 1 {
			if xDiff > 0 {
				// Moving to wards the right x increases
				// The +1 is to make sure that we do not check the starting square
				if e.Board[startingY][(startingX+1)+colOffset] != nil {
					return fmt.Errorf("move not allowed, rook path has a piece blocking i.e (%v,%v)", startingX+1+colOffset, startingY)
				}
			} else {
				// Moving to wards the left x decreases
				// The -1 is to make sure that we do not check the starting square
				if e.Board[startingY][(startingX-1)-colOffset] != nil {
					return fmt.Errorf("move not allowed, rook path has a piece blocking i.e (%v,%v)", startingX+1+colOffset, startingY)
				}
			}
		}
	} else if yAd > xAd {
		for rowOffset := range yAd - 1 {
			if yDiff > 0 {
				if e.Board[(startingY+1)+rowOffset][startingX] != nil {
					return fmt.Errorf("move not allowed, rook path has a piece blocking i.e (%v,%v)", startingX, startingY+1+rowOffset)
				}
			} else {
				if e.Board[(startingY-1)-rowOffset][startingX] != nil {
					return fmt.Errorf("move not allowed, rook path has a piece blocking i.e (%v,%v)", startingX, startingY+1+rowOffset)
				}
			}
		}
	}

	e.Board[destinationY][destinationX] = e.Board[startingY][startingX]
	e.Board[startingY][startingX] = nil
	return nil
}

func (e *Engine) MoveQueen() error {
	return nil
}
