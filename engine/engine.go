// Package engine - has all the internal working of the game. That includes:
//  1. Pieces
//  2. Board logic
//  3. Move validation
//  4. Check and checkmate detection
package engine

import "log"

type Engine struct {
	Board              [8][8]*Piece
	CurrentPlayerColor PieceColor
	EnpassantSquare    *Square
	castleRights       CastleRights
	MoveStack          []Move
}

func NewEngine() *Engine {
	engine := &Engine{
		CurrentPlayerColor: White,
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
		e.Board[1][col] = &Piece{
			Type:  Pawn,
			Color: White,
		}
		e.Board[6][col] = &Piece{
			Type:  Pawn,
			Color: Black,
		}
	}

	// Other pieces
	oderOfPieces := [8]PieceType{Rook, Knight, Bishop, Queen, King, Bishop, Knight, Rook}
	for col := range 8 {
		piece := oderOfPieces[col]
		e.Board[0][col] = &Piece{
			Type:  piece,
			Color: White,
		}
		e.Board[7][col] = &Piece{
			Type:  piece,
			Color: Black,
		}
	}
}

func (e *Engine) String() string {
	var output string
	for row := range 8 {
		for col := range 8 {
			piece := e.Board[row][col]
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
