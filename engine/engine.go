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
}

func NewEngine() *Engine {
	engine := &Engine{
		CurrentPlayerColor: White,
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
			Type:         Pawn,
			Color:        White,
			RenderLetter: "P",
		}
		e.Board[6][col] = &Piece{
			Type:         Pawn,
			Color:        Black,
			RenderLetter: "P",
		}
	}

	// Other pieces
	oderOfPieces := [8]PieceType{Rook, Knight, Bishop, King, Queen, Bishop, Knight, Rook}
	for col := range 8 {
		piece := oderOfPieces[col]
		renderLetter := GetRenderLetter(piece)
		e.Board[0][col] = &Piece{
			Type:         piece,
			Color:        White,
			RenderLetter: renderLetter,
		}
		e.Board[7][col] = &Piece{
			Type:         piece,
			Color:        Black,
			RenderLetter: renderLetter,
		}
	}
}
