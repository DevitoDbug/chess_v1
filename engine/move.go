package engine

import (
	"fmt"
)

func (e *Engine) MovePiece(input Input) error {
	piece := e.board[input.StartY][input.StartX]
	if piece == nil {
		return fmt.Errorf("no piece in the square referenced")
	}

	// Catch players moving other opponents pieces early
	if piece.Color != e.currentPlayerColor {
		return fmt.Errorf("invalid move, you can only move pieces belonging to the current player")
	}

	move, err := e.GigaMove(input.StartX, input.StartY, input.DestinationX, input.DestinationY)
	if err != nil {
		return err
	}

	playerIsStillInCheck := e.isCurrentPlayersKingInCheck()
	if playerIsStillInCheck {
		err := e.UndoMove(move)
		if err != nil {
			return fmt.Errorf("could not undo move. Error:  %v", err)
		}

		return fmt.Errorf("invalid move, player in check")
	}

	e.moveStack = append(e.moveStack, move)

	// Clean ups
	if e.enpassantSquare != nil {
		if e.enpassantSquare.EnpassantChecked {
			e.enpassantSquare = nil
		} else {
			e.enpassantSquare.EnpassantChecked = true
		}
	}

	return nil
}
