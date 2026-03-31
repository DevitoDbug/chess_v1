package engine

import (
	"fmt"
)

func (e *Engine) MovePiece(input Input) error {
	startPiece := e.board[input.StartY][input.StartX]
	if startPiece == nil {
		return fmt.Errorf("no piece in the square referenced")
	}

	// Catch players moving other opponents pieces early
	if startPiece.Color != e.currentPlayerColor {
		return fmt.Errorf("invalid move, you can only move pieces belonging to the current player")
	}

	move, err := e.GigaMove(input.StartX, input.StartY, input.DestinationX, input.DestinationY)
	if err != nil {
		return err
	}

	playerIsStillInCheck := e.isCurrentPlayersKingInCheck(startPiece.Color)
	if playerIsStillInCheck {
		err := e.UndoMove(move)
		if err != nil {
			return fmt.Errorf("could not undo move. Error:  %v", err)
		}

		return fmt.Errorf("invalid move, player in check")
	}

	// Check for promotions
	movedPiece := e.board[move.ToY][move.ToX]
	if movedPiece.Color == White && input.DestinationY == 7 && movedPiece.Type == Pawn {
		if input.PromotionPiece != nil {
			e.board[input.DestinationY][input.DestinationX].Type = *input.PromotionPiece
		} else {
			err := e.UndoMove(move)
			if err != nil {
				return fmt.Errorf("could not undo move. Error:  %v", err)
			}
			return fmt.Errorf("invalid move, promotion piece not specified")
		}
	}
	if movedPiece.Color == Black && input.DestinationY == 0 && movedPiece.Type == Pawn {
		if input.PromotionPiece != nil {
			e.board[input.DestinationY][input.DestinationX].Type = *input.PromotionPiece
		} else {
			err := e.UndoMove(move)
			if err != nil {
				return fmt.Errorf("could not undo move. Error:  %v", err)
			}
			return fmt.Errorf("invalid move, promotion piece not specified")
		}
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
