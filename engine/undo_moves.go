package engine

import "fmt"

func (e *Engine) UndoMove(move Move) error {
	if e.board[move.FromY][move.FromX] != nil || e.board[move.ToY][move.ToX] == nil {
		return fmt.Errorf("piece not found for either the destination or the start insinuated by the move claim")
	}
	pieceMoved := e.board[move.ToY][move.ToX]
	e.board[move.FromY][move.FromX] = pieceMoved
	e.board[move.ToY][move.ToX] = nil

	// Check if it was an enpassant if it was we need to restore the piece that was taken out  of the board
	if move.IsEnpassant && move.CapturedPiece != nil {
		var y int32
		switch pieceMoved.Color {
		case White:
			y = move.ToY - 1
		case Black:
			y = move.ToY + 1
		}
		e.board[y][move.ToX] = move.CapturedPiece

	} else if move.CapturedPiece != nil {
		e.board[move.ToY][move.ToX] = move.CapturedPiece
	}

	if move.IsPromotion {
		e.board[move.FromY][move.FromX].Type = Pawn
	}

	if move.IsCastling {
		var rookPosX, rookY, rookOrigPosX int32
		switch pieceMoved.Color {
		case White:
			rookY = int32(0)
		case Black:
			rookY = int32(7)
		}

		switch move.ToX {
		case 2:
			rookPosX = move.ToX + 1
			rookOrigPosX = 0
		case 6:
			rookPosX = move.ToX - 1
			rookOrigPosX = 7
		}

		// Confirm rook is at the position should be after a castling happens
		rook := e.board[rookY][rookPosX]
		if rook == nil {
			return fmt.Errorf("invalid undo move, king castling did not place rook in the right position")
		}

		e.board[rookY][rookOrigPosX] = rook
		e.board[rookY][rookPosX] = nil
	}

	e.enpassantSquare = move.PreviousEnpassantSquareState
	e.castleRights = move.PreviousCastlingState
	return nil
}
