package engine

type PieceType int

const (
	Pawn PieceType = iota
	Knight
	Bishop
	Rook
	Queen
	King
)

type PieceColor string

const (
	Black PieceColor = "black"
	White PieceColor = "white"
)

type Piece struct {
	Type  PieceType
	Color PieceColor
}

const (
	RowNumber    = 8
	ColumnNumber = 8
)

type Square struct {
	RowIndex    int32
	ColumnIndex int32
	Checked     bool
}

type CastleRights struct {
	WhiteKingSideCastle  bool
	WhiteQueenSideCastle bool
	BlackKingSideCastle  bool
	BlackQueenSideCastle bool
}

type Move struct {
	FromX int32
	FromY int32
	ToX   int32
	ToY   int32

	CapturedPiece *Piece

	IsEnpassant  bool
	IsCastling   bool
	IsPromortion bool
}
