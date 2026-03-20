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

func (p PieceType) String() string {
	switch p {
	case Pawn:
		return "Pawn"
	case Knight:
		return "Knight"
	case Bishop:
		return "Bishop"
	case Rook:
		return "Rook"
	case Queen:
		return "Queen"
	case King:
		return "King"
	default:
		return "Unknown"
	}
}

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
	RowIndex         int32
	ColumnIndex      int32
	EnpassantChecked bool // Used mainly for enpassant ( has this square ever been checked for enpassant )
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

	CapturedPiece                *Piece
	PreviousEnpassantSquareState *Square
	PreviousCastlingState        CastleRights

	IsEnpassant bool
	IsCastling  bool
	IsPromotion bool
}

type EndGameState struct {
	StaleMate bool
	CheckMate bool
}
