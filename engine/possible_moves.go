package engine

import (
	"fmt"
)

// Have functions that will produce possible moves for all pieces based on the board
// Sliding moves will probably be the tricky ones
func (e *Engine) GetPossibleMoves(sqx, sqy int32) []Move {
	moves := []Move{}
	piece := e.board[sqy][sqx]
	if piece == nil {
		return nil
	}

	switch piece.Type {
	case Pawn:
		moves = e.allPossiblePawnMoves(sqx, sqy)
	case King:
		moves = e.allPossibleKingMoves(sqx, sqy)
	case Queen:
		fmt.Print("hello")
	case Rook:
		fmt.Print("hello")
	case Bishop:
		fmt.Print("hello")
	case Knight:
		fmt.Print("hello")
	default:
		return nil
	}

	return moves
}

func (e *Engine) allPossiblePawnMoves(sqx, sqy int32) []Move {
	moves := []Move{}
	pawn := e.board[sqy][sqx]
	if pawn == nil {
		return nil
	}
	var dy, homeRow, promotionRow int32

	switch pawn.Color {
	case White:
		dy = 1
		homeRow = 1
		promotionRow = 7
	case Black:
		dy = -1
		homeRow = 6
		promotionRow = 0
	}

	// Forward move should be possible
	if e.isInsideBoard(sqx, sqy+dy) && e.board[sqy+dy][sqx] == nil {
		isPromotion := false
		if sqy+dy == promotionRow {
			isPromotion = true
		}

		moves = append(moves, Move{
			FromX: sqx,
			FromY: sqy,
			ToX:   sqx,
			ToY:   (sqy + dy),

			IsPromotion: isPromotion,
		})

		if e.isInsideBoard(sqx, sqy+(2*dy)) && e.board[sqy+(2*dy)][sqx] == nil && sqy == homeRow {
			// Double forward move from the pawn home square should be possible
			moves = append(moves, Move{
				FromX: sqx,
				FromY: sqy,
				ToX:   sqx,
				ToY:   (sqy + (2 * dy)),
			})
		}
	}

	// Pawn takes move should be possible in both right and left
	if e.isInsideBoard(sqx+1, sqy+dy) {
		isPromotion := false
		if sqy+dy == promotionRow {
			isPromotion = true
		}

		pieceInRight := e.board[sqy+dy][sqx+1]
		if pieceInRight != nil && pieceInRight.Color != pawn.Color {
			moves = append(moves, Move{
				FromX: sqx,
				FromY: sqy,
				ToX:   sqx + 1,
				ToY:   sqy + dy,

				CapturedPiece: e.board[sqy+dy][sqx+1],
				IsPromotion:   isPromotion,
			})
		}
	}

	if e.isInsideBoard(sqx-1, sqy+dy) {
		isPromotion := false
		if sqy+dy == promotionRow {
			isPromotion = true
		}

		pieceInLeft := e.board[sqy+dy][sqx-1]
		if pieceInLeft != nil && pieceInLeft.Color != pawn.Color {
			moves = append(moves, Move{
				FromX: sqx,
				FromY: sqy,
				ToX:   sqx - 1,
				ToY:   sqy + dy,

				CapturedPiece: e.board[sqy+dy][sqx-1],
				IsPromotion:   isPromotion,
			})
		}
	}

	// Enpassant move
	if e.enpassantSquare != nil {
		if sqy+dy == e.enpassantSquare.RowIndex && sqx+1 == e.enpassantSquare.ColumnIndex {
			// Enpassant to the right
			moves = append(moves, Move{
				FromX: sqx,
				FromY: sqy,
				ToX:   sqx + 1,
				ToY:   sqy + dy,

				IsEnpassant: true,
			})
		} else if sqy+dy == e.enpassantSquare.RowIndex && sqx-1 == e.enpassantSquare.ColumnIndex {
			// Enpassant to the left
			moves = append(moves, Move{
				FromX: sqx,
				FromY: sqy,
				ToX:   sqx - 1,
				ToY:   sqy + dy,

				IsEnpassant: true,
			})
		}
	}

	return moves
}

func (e *Engine) allPossibleKingMoves(sqx, sqy int32) []Move {
	moves := []Move{}
	possibilities := [8][2]int32{
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
		{-1, -1},
		{-1, 1},
		{-1, 0},
		{0, 1},
	}

	kingSquare := e.board[sqy][sqx]
	if kingSquare == nil || kingSquare.Type != King {
		return moves
	}
	opponentColor := toggleCurrentPlayer(kingSquare.Color)

	for _, d := range possibilities {
		destinationY := sqy + d[0]
		destinationX := sqx + d[1]
		if e.isInsideBoard(destinationX, destinationY) && !e.isSquareAttacked(destinationX, destinationY, opponentColor) {
			destinationPiece := e.board[destinationY][destinationX]
			if destinationPiece == nil {
				moves = append(moves, Move{
					FromX:                  sqx,
					FromY:                  sqy,
					ToX:                    destinationX,
					ToY:                    destinationY,
					PreviousCastlingState:  e.castleRights,
					PreviousEnpassantState: e.enpassantSquare,
				})
			} else if destinationPiece.Color == opponentColor {
				moves = append(moves, Move{
					FromX: sqx,
					FromY: sqy,
					ToX:   destinationX,
					ToY:   destinationY,
					CapturedPiece: &Piece{
						Type:  destinationPiece.Type,
						Color: opponentColor,
					},
					PreviousCastlingState:  e.castleRights,
					PreviousEnpassantState: e.enpassantSquare,
				})
			}
		}
	}

	if sqx == 4 && !e.isSquareAttacked(sqx, sqy, opponentColor) {
		if kingSquare.Color == White && sqy == 0 {
			if e.castleRights.WhiteKingSideCastle &&
				!e.isSquareAttacked(5, 0, opponentColor) && e.board[0][5] == nil &&
				!e.isSquareAttacked(6, 0, opponentColor) && e.board[0][6] == nil {
				moves = append(moves, Move{
					FromX:                  sqx,
					FromY:                  sqy,
					ToX:                    6,
					ToY:                    sqy,
					PreviousEnpassantState: e.enpassantSquare,
					PreviousCastlingState:  e.castleRights,
					IsCastling:             true,
				})
			}
			if e.castleRights.WhiteQueenSideCastle &&
				!e.isSquareAttacked(1, 0, opponentColor) && e.board[0][1] == nil &&
				!e.isSquareAttacked(2, 0, opponentColor) && e.board[0][2] == nil &&
				!e.isSquareAttacked(3, 0, opponentColor) && e.board[0][3] == nil {
				moves = append(moves, Move{
					FromX:                  sqx,
					FromY:                  sqy,
					ToX:                    2,
					ToY:                    sqy,
					PreviousEnpassantState: e.enpassantSquare,
					PreviousCastlingState:  e.castleRights,
					IsCastling:             true,
				})
			}
		} else if kingSquare.Color == Black && sqy == 7 {
			if e.castleRights.BlackKingSideCastle &&
				!e.isSquareAttacked(5, 7, opponentColor) && e.board[7][5] == nil &&
				!e.isSquareAttacked(6, 7, opponentColor) && e.board[7][6] == nil {
				moves = append(moves, Move{
					FromX:                  sqx,
					FromY:                  sqy,
					ToX:                    6,
					ToY:                    sqy,
					PreviousEnpassantState: e.enpassantSquare,
					PreviousCastlingState:  e.castleRights,
					IsCastling:             true,
				})
			}
			if e.castleRights.BlackQueenSideCastle &&
				!e.isSquareAttacked(1, 7, opponentColor) && e.board[7][1] == nil &&
				!e.isSquareAttacked(2, 7, opponentColor) && e.board[7][2] == nil &&
				!e.isSquareAttacked(3, 7, opponentColor) && e.board[7][3] == nil {
				moves = append(moves, Move{
					FromX:                  sqx,
					FromY:                  sqy,
					ToX:                    2,
					ToY:                    sqy,
					PreviousEnpassantState: e.enpassantSquare,
					PreviousCastlingState:  e.castleRights,
					IsCastling:             true,
				})
			}
		}
	}

	return moves
}
