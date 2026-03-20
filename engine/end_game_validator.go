package engine

import "fmt"

// GetEndGameState - tells if there is a checkmate or a stalemate as early as possible.
func (e *Engine) GetEndGameState() *EndGameState {
	fmt.Println("Endgame state called")
	opponentColor := toggleCurrentPlayer(e.currentPlayerColor)

	for row := range int32(8) {
		for column := range int32(8) {
			piece := e.board[row][column]

			if piece == nil || piece.Color != opponentColor {
				continue
			}

			pseudoMoves := e.GetAllPossiblePseudoMoves(column, row)
			for _, pseudoMove := range pseudoMoves {
				kingIsInDanger := false
				move, err := e.GigaMove(pseudoMove.FromX, pseudoMove.FromY, pseudoMove.ToX, pseudoMove.ToY)
				if err != nil {
					continue
				}

				king := e.findKing(opponentColor)
				if e.isSquareAttacked(king.ColumnIndex, king.RowIndex, e.currentPlayerColor) {
					kingIsInDanger = true
				}

				err = e.UndoMove(move)
				if err != nil {
					panic("UndoMove failed")
				}

				if kingIsInDanger {
					fmt.Println("King is danger for move: ")
					fmt.Printf("%+v\n", move)
					continue
				}

				fmt.Println("Move found")
				fmt.Printf("%+v\n", move)
				fmt.Println()

				return nil // found a valid move no need for further checks
			}
		}
	}

	king := e.findKing(opponentColor)
	if king == nil {
		panic("King not found ")
	}

	if e.isSquareAttacked(king.ColumnIndex, king.RowIndex, e.currentPlayerColor) {
		return &EndGameState{
			StaleMate: false,
			CheckMate: true,
		}
	}

	// find out if it is a stalemate or a checkmate
	return &EndGameState{
		StaleMate: true,
		CheckMate: false,
	}
}

func (e *Engine) findKing(color PieceColor) *Square {
	// Iterate through the board and find the king
	for row := range int32(8) {
		for col := range int32(8) {
			piece := e.board[row][col]

			if piece != nil && piece.Type == King && piece.Color == color {
				return &Square{
					RowIndex:    row,
					ColumnIndex: col,
				}
			}
		}
	}

	return nil
}
