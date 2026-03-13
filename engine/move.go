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
		err := e.MoveKing(input.StartX, input.StartY, input.DestinationX, input.DestinationY)
		if err != nil {
			return err
		}
	case Queen:
		err := e.MoveQueen(input.StartX, input.StartY, input.DestinationX, input.DestinationY)
		if err != nil {
			return err
		}
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
		err := e.MoveKnight(input.StartX, input.StartY, input.DestinationX, input.DestinationY)
		if err != nil {
			return err
		}
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

func (e *Engine) MovePawn(startX, startY, destinationX, destinationY int32) error {
	startPiece := e.Board[startY][startX]
	destinationPiece := e.Board[destinationY][destinationX]

	x1 := startX
	y1 := startY
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

		e.Board[destinationY][destinationX] = e.Board[startY][startX]
		e.Board[startY][startX] = nil
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
		switch startPiece.Color {
		// Moving two steps creates possibility for ampersand in the position (start position + 1) or
		// (destination position -1 ) from whites perspective
		case White:
			if startY != 1 {
				return fmt.Errorf("paws can only move two moves if it is the first move")
			}
			e.EnpassantSquare = &Square{
				RowIndex:    int32(startY + 1),
				ColumnIndex: startX,
			}
		case Black:
			if startY != 6 {
				return fmt.Errorf("paws can only move two moves if it is the first move")
			}
			e.EnpassantSquare = &Square{
				RowIndex:    int32(startY - 1),
				ColumnIndex: startX,
			}
		default:
			return fmt.Errorf("pawn color not recognized")
		}
	}

	if e.Board[destinationY][destinationX] != nil {
		return fmt.Errorf("pawns cannot move to squares that have pieces")
	}

	e.Board[destinationY][destinationX] = e.Board[startY][startX]
	e.Board[startY][startX] = nil

	return nil
}

func (e *Engine) MoveKing(startX, startY, destinationX, destinationY int32) error {
	piece := e.Board[startY][startX]
	if piece == nil || piece.Type != King {
		return fmt.Errorf("no king at starting square")
	}
	color := e.Board[startY][startX].Color

	if e.isCastlingMove(startX, startY, destinationX, destinationY) {
		// Move the king
		e.Board[destinationY][destinationX] = e.Board[startY][startX]
		e.Board[startY][startX] = nil

		// Move rook
		if destinationX > startX { // King-side
			e.Board[startY][destinationX-1] = e.Board[startY][7]
			e.Board[startY][7] = nil
		} else { // Queen-side
			e.Board[startY][destinationX+1] = e.Board[startY][0]
			e.Board[startY][0] = nil
		}
	} else {
		if err := e.isValidKingMove(startX, startY, destinationX, destinationY); err != nil {
			return err
		}
		e.Board[destinationY][destinationX] = e.Board[startY][startX]
		e.Board[startY][startX] = nil
	}
	// Disable castling for the corresponding color side
	switch color {
	case White:
		e.castleRights.WhiteKingSideCastle = false
		e.castleRights.WhiteQueenSideCastle = false
	case Black:
		e.castleRights.BlackKingSideCastle = false
		e.castleRights.BlackQueenSideCastle = false
	}

	return nil
}

func (e *Engine) MoveKnight(startX, startY, destinationX, destinationY int32) error {
	if e.Board[destinationY][destinationX] != nil && e.Board[destinationY][destinationX].Color == e.CurrentPlayerColor {
		return fmt.Errorf("move not allowed. Players cannot attack their own piece")
	}

	x1 := startX
	x2 := destinationX
	y1 := startY
	y2 := destinationY

	xAb := utils.AbsoluteDiff(x1, x2)
	yAb := utils.AbsoluteDiff(y1, y2)

	// One of the values must be a 1 and the other a 2
	if xAb*yAb != 2 {
		return fmt.Errorf("move not allowed")
	}

	e.Board[destinationY][destinationX] = e.Board[startY][startX]
	e.Board[startY][startX] = nil

	return nil
}

func (e *Engine) MoveBishop(startX, startY, destinationX, destinationY int32) error {
	err := e.isDiagonalValidSlidingMove(startX, startY, destinationX, destinationY)
	if err != nil {
		return err
	}

	e.Board[destinationY][destinationX] = e.Board[startY][startX]
	e.Board[startY][startX] = nil
	return nil
}

func (e *Engine) MoveRook(startX, startY, destinationX, destinationY int32) error {
	err := e.isStraightValidSlidingMove(startX, startY, destinationX, destinationY)
	if err != nil {
		return err
	}

	e.Board[destinationY][destinationX] = e.Board[startY][startX]
	e.Board[startY][startX] = nil

	// Rook move from home square should warrant a casting right elimination of that side of the king
	if startX == 0 && startY == 0 {
		e.castleRights.WhiteQueenSideCastle = false
	} else if startX == 7 && startY == 0 {
		e.castleRights.WhiteKingSideCastle = false
	} else if startX == 7 && startY == 7 {
		e.castleRights.BlackKingSideCastle = false
	} else if startX == 0 && startY == 7 {
		e.castleRights.BlackQueenSideCastle = false
	}
	return nil
}

func (e *Engine) MoveQueen(startX, startY, destinationX, destinationY int32) error {
	// A queen move is just rook + bishop moves
	if err := e.isDiagonalValidSlidingMove(startX, startY, destinationX, destinationY); err != nil {
		err = e.isStraightValidSlidingMove(startX, startY, destinationX, destinationY)
		if err != nil {
			return fmt.Errorf("move not allowed")
		}
	}

	e.Board[destinationY][destinationX] = e.Board[startY][startX]
	e.Board[startY][startX] = nil
	return nil
}
