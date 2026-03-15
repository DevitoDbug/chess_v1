// Package engine - has all the internal working of the game. That includes:
//  1. Pieces
//  2. Board logic
//  3. Move validation
//  4. Check and checkmate detection
package engine

import "log"

type Engine struct {
	board              [8][8]*Piece
	currentPlayerColor PieceColor
	enpassantSquare    *Square
	castleRights       CastleRights
	moveStack          []Move
}

func NewEngine() *Engine {
	engine := &Engine{
		currentPlayerColor: White,
		castleRights: CastleRights{
			WhiteKingSideCastle:  true,
			WhiteQueenSideCastle: true,
			BlackKingSideCastle:  true,
			BlackQueenSideCastle: true,
		},
	}

	engine.Init() // Fucking annoying to export initializations
	return engine
}

func (e *Engine) Init() {
	if e == nil {
		log.Fatal("Engine not set")
		return
	}

	// Pawns
	for col := range 8 {
		e.board[1][col] = &Piece{
			Type:  Pawn,
			Color: White,
		}
		e.board[6][col] = &Piece{
			Type:  Pawn,
			Color: Black,
		}
	}

	// Other pieces
	oderOfPieces := [8]PieceType{Rook, Knight, Bishop, Queen, King, Bishop, Knight, Rook}
	for col := range 8 {
		piece := oderOfPieces[col]
		e.board[0][col] = &Piece{
			Type:  piece,
			Color: White,
		}
		e.board[7][col] = &Piece{
			Type:  piece,
			Color: Black,
		}
	}
}

func (e *Engine) String() string {
	var output string
	for row := range 8 {
		for col := range 8 {
			piece := e.board[row][col]
			if piece != nil {
				output += string(GetRenderLetter(piece.Type, piece.Color))
			} else {
				output += "0"
			}
		}
		output += "\n"
	}

	return output
}
