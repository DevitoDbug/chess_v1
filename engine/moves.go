package engine

import (
	"fmt"

	"github.com/DevitoDbug/chess_v1/utils"
)

func (e *Engine) GigaMove(startX, startY, destinationX, destinationY int32) (Move, error) {
	move := Move{}
	fmt.Printf("startX: %v\nstartY: %v\n", startX, startY)

	if !e.isInsideBoard(startX, startY) {
		return move, fmt.Errorf("invalid move, referenced address outside the board")
	}

	startingPiece := e.board[startY][startX]
	if startingPiece == nil {
		return move, fmt.Errorf("invalid move, no piece in starting square")
	}

	switch startingPiece.Type {
	case Pawn:
		return e.MovePawn(startX, startY, destinationX, destinationY)
	case King:
		return e.MoveKing(startX, startY, destinationX, destinationY)
	case Queen:
		return e.MoveQueen(startX, startY, destinationX, destinationY)
	case Knight:
		return e.MoveKnight(startX, startY, destinationX, destinationY)
	case Bishop:
		return e.MoveBishop(startX, startY, destinationX, destinationY)
	case Rook:
		return e.MoveRook(startX, startY, destinationX, destinationY)
	}

	return move, fmt.Errorf("invalid move, where the fuck did you get a piece we have never seen")
}

func (e *Engine) MovePawn(startX, startY, destinationX, destinationY int32) (Move, error) {
	move := Move{}
	move.PreviousEnpassantSquareState = e.enpassantSquare
	move.PreviousCastlingState = e.castleRights

	startPiece := e.board[startY][startX]
	currentPlayerColor := startPiece.Color
	destinationPiece := e.board[destinationY][destinationX]

	x1 := startX
	y1 := startY
	x2 := destinationX
	y2 := destinationY
	xAbsDiff := utils.AbsoluteDiff(x1, x2)
	yAbsDiff := utils.AbsoluteDiff(y1, y2)

	xDiff := x2 - x1
	yDiff := int32(0)
	switch currentPlayerColor {
	// White moves are positive while black moves are negative
	// White moves from a 0 -> 7 while black moves from 7 -> 0
	case White:
		yDiff = y2 - y1
	case Black:
		yDiff = y1 - y2
	}

	// Capture piece or enpassant
	if xAbsDiff == 1 && yAbsDiff == 1 && yDiff > 0 {
		// yDiff check is to make sure that we are moving in the right direction,
		// white is not moving toward 0 and black towards 7
		isEnpassantMove := e.enpassantSquare != nil && destinationY == e.enpassantSquare.RowIndex && destinationX == e.enpassantSquare.ColumnIndex
		isAttackingOpponentPiece := destinationPiece != nil && destinationPiece.Color != currentPlayerColor
		if !isAttackingOpponentPiece && !isEnpassantMove {
			return move, fmt.Errorf("player is not allowed to attack his own piece")
		}

		if isEnpassantMove {
			move.IsEnpassant = true

			switch currentPlayerColor {
			case White:
				move.CapturedPiece = e.board[e.enpassantSquare.RowIndex-1][e.enpassantSquare.ColumnIndex]
				e.board[e.enpassantSquare.RowIndex-1][e.enpassantSquare.ColumnIndex] = nil
			case Black:
				move.CapturedPiece = e.board[e.enpassantSquare.RowIndex+1][e.enpassantSquare.ColumnIndex]
				e.board[e.enpassantSquare.RowIndex+1][e.enpassantSquare.ColumnIndex] = nil
			}
		}

		e.board[destinationY][destinationX] = e.board[startY][startX]
		e.board[startY][startX] = nil
		move.FromX = startX
		move.FromY = startY
		move.ToX = destinationX
		move.ToY = destinationY
		return move, nil
	}

	if xDiff != 0 {
		return move, fmt.Errorf("pawns are only allowed to move forward or attack diagonally one step")
	}

	if yDiff == 0 {
		return move, fmt.Errorf("pawn did not move")
	}

	if yDiff != 2 && yDiff != 1 {
		return move, fmt.Errorf("pawns are only allowed to move one step forward, or two steps for the first move")
	}

	if yDiff == 2 {
		switch startPiece.Color {
		// Moving two steps creates possibility for enpassant in the position (start position + 1) or
		// (destination position -1 ) from whites perspective
		case White:
			if startY != 1 {
				return move, fmt.Errorf("paws can only move two moves if it is the first move")
			}
			e.enpassantSquare = &Square{
				RowIndex:    int32(startY + 1),
				ColumnIndex: startX,
			}
		case Black:
			if startY != 6 {
				return move, fmt.Errorf("paws can only move two moves if it is the first move")
			}
			e.enpassantSquare = &Square{
				RowIndex:    int32(startY - 1),
				ColumnIndex: startX,
			}
		default:
			return move, fmt.Errorf("pawn color not recognized")
		}
	}

	if e.board[destinationY][destinationX] != nil {
		return move, fmt.Errorf("pawns cannot move to squares that have pieces")
	}

	e.board[destinationY][destinationX] = e.board[startY][startX]
	e.board[startY][startX] = nil
	move.FromX = startX
	move.FromY = startY
	move.ToX = destinationX
	move.ToY = destinationY

	return move, nil
}

func (e *Engine) MoveKing(startX, startY, destinationX, destinationY int32) (Move, error) {
	move := Move{}
	move.PreviousEnpassantSquareState = e.enpassantSquare
	move.PreviousCastlingState = e.castleRights

	destinationPiece := e.board[destinationY][destinationX]
	startPiece := e.board[startY][startX]
	if startPiece == nil || startPiece.Type != King {
		return move, fmt.Errorf("no king at starting square")
	}
	color := e.board[startY][startX].Color

	if e.isCastlingMove(startX, startY, destinationX, destinationY) {
		// Move the king
		e.board[destinationY][destinationX] = e.board[startY][startX]
		e.board[startY][startX] = nil

		// Move rook
		if destinationX > startX { // King-side
			e.board[startY][destinationX-1] = e.board[startY][7]
			e.board[startY][7] = nil
		} else { // Queen-side
			e.board[startY][destinationX+1] = e.board[startY][0]
			e.board[startY][0] = nil
		}
		move.IsCastling = true
	} else {
		if err := e.isValidKingMove(startX, startY, destinationX, destinationY); err != nil {
			return move, err
		}
		e.board[destinationY][destinationX] = e.board[startY][startX]
		e.board[startY][startX] = nil
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
	move.FromX = startX
	move.FromY = startY
	move.ToX = destinationX
	move.ToY = destinationY
	move.CapturedPiece = destinationPiece
	return move, nil
}

func (e *Engine) MoveQueen(startX, startY, destinationX, destinationY int32) (Move, error) {
	move := Move{}
	move.PreviousEnpassantSquareState = e.enpassantSquare
	move.PreviousCastlingState = e.castleRights

	destinationPiece := e.board[destinationY][destinationX]

	// A queen move is just rook + bishop moves
	if err := e.isDiagonalValidSlidingMove(startX, startY, destinationX, destinationY); err != nil {
		err = e.isStraightValidSlidingMove(startX, startY, destinationX, destinationY)
		if err != nil {
			return move, fmt.Errorf("move not allowed")
		}
	}

	e.board[destinationY][destinationX] = e.board[startY][startX]
	e.board[startY][startX] = nil
	move.FromX = startX
	move.FromY = startY
	move.ToX = destinationX
	move.ToY = destinationY
	move.CapturedPiece = destinationPiece

	return move, nil
}

func (e *Engine) MoveKnight(startX, startY, destinationX, destinationY int32) (Move, error) {
	move := Move{}
	move.PreviousEnpassantSquareState = e.enpassantSquare
	move.PreviousCastlingState = e.castleRights

	startPiece := e.board[startY][startX]
	currentPlayerColor := startPiece.Color
	destinationPiece := e.board[destinationY][destinationX]
	if destinationPiece != nil && destinationPiece.Color == currentPlayerColor {
		return move, fmt.Errorf("move not allowed. Players cannot attack their own piece")
	}

	x1 := startX
	x2 := destinationX
	y1 := startY
	y2 := destinationY

	xAb := utils.AbsoluteDiff(x1, x2)
	yAb := utils.AbsoluteDiff(y1, y2)

	// One of the values must be a 1 and the other a 2
	if xAb*yAb != 2 {
		return move, fmt.Errorf("move not allowed")
	}

	e.board[destinationY][destinationX] = e.board[startY][startX]
	e.board[startY][startX] = nil
	move.FromX = startX
	move.FromY = startY
	move.ToX = destinationX
	move.ToY = destinationY
	move.CapturedPiece = destinationPiece

	return move, nil
}

func (e *Engine) MoveBishop(startX, startY, destinationX, destinationY int32) (Move, error) {
	move := Move{}
	move.PreviousEnpassantSquareState = e.enpassantSquare
	move.PreviousCastlingState = e.castleRights

	destinationPiece := e.board[destinationY][destinationX]

	err := e.isDiagonalValidSlidingMove(startX, startY, destinationX, destinationY)
	if err != nil {
		return move, err
	}

	e.board[destinationY][destinationX] = e.board[startY][startX]
	e.board[startY][startX] = nil
	move.FromX = startX
	move.FromY = startY
	move.ToX = destinationX
	move.ToY = destinationY
	move.CapturedPiece = destinationPiece

	return move, nil
}

func (e *Engine) MoveRook(startX, startY, destinationX, destinationY int32) (Move, error) {
	move := Move{}
	move.PreviousEnpassantSquareState = e.enpassantSquare
	move.PreviousCastlingState = e.castleRights

	destinationPiece := e.board[destinationY][destinationX]

	err := e.isStraightValidSlidingMove(startX, startY, destinationX, destinationY)
	if err != nil {
		return move, err
	}

	e.board[destinationY][destinationX] = e.board[startY][startX]
	e.board[startY][startX] = nil

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

	move.FromX = startX
	move.FromY = startY
	move.ToX = destinationX
	move.ToY = destinationY
	move.CapturedPiece = destinationPiece
	return move, nil
}
