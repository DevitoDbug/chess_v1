package engine

func GetRenderLetter(pieceType PieceType) string {
	switch pieceType {
	case Pawn:
		return "P"
	case Knight:
		return "N"
	case Bishop:
		return "B"
	case Rook:
		return "R"
	case Queen:
		return "Q"
	case King:
		return "K"
	default:
		return ""
	}
}

func GetColorLetter(color PieceColor) string {
	switch color {
	case Black:
		return "B"
	case White:
		return "W"
	default:
		return " "
	}
}
