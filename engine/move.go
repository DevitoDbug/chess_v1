package engine

import (
	"fmt"
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
		move, err := e.MovePawn(input.StartX, input.StartY, input.DestinationX, input.DestinationY)
		if err != nil {
			return err
		}
		// TODO: Implement the undo for the cases of king in check
		e.MoveStack = append(e.MoveStack, move)
	case King:
		move, err := e.MoveKing(input.StartX, input.StartY, input.DestinationX, input.DestinationY)
		if err != nil {
			return err
		}
		e.MoveStack = append(e.MoveStack, move)
	case Queen:
		move, err := e.MoveQueen(input.StartX, input.StartY, input.DestinationX, input.DestinationY)
		if err != nil {
			return err
		}
		e.MoveStack = append(e.MoveStack, move)
	case Rook:
		move, err := e.MoveRook(input.StartX, input.StartY, input.DestinationX, input.DestinationY)
		if err != nil {
			return err
		}
		e.MoveStack = append(e.MoveStack, move)
	case Bishop:
		move, err := e.MoveBishop(input.StartX, input.StartY, input.DestinationX, input.DestinationY)
		if err != nil {
			return err
		}
		e.MoveStack = append(e.MoveStack, move)
	case Knight:
		move, err := e.MoveKnight(input.StartX, input.StartY, input.DestinationX, input.DestinationY)
		if err != nil {
			return err
		}
		e.MoveStack = append(e.MoveStack, move)
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
