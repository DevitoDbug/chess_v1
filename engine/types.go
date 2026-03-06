package engine

type PieceType string

const (
	Pawn   PieceType = "pawn"
	Knight PieceType = "knight"
	Bishop PieceType = "piece"
	Rook   PieceType = "rook"
	Queen  PieceType = "queen"
	King   PieceType = "king"
)

type PieceColor string

const (
	Black PieceColor = "black"
	White PieceColor = "white"
)

type Piece struct {
	Type         PieceType
	Color        PieceColor
	RenderLetter string
}

const (
	RowNumber    = 8
	ColumnNumber = 8
)
