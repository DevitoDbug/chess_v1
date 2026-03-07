package engine

import (
	"fmt"
	"strings"
)

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

type Input struct {
	Color        PieceColor
	Type         PieceType
	DestinationX int
	DestinationY int
	StartX       int
	StartY       int
	Piece        *Piece
}

// Expected sample input BN-a2
func (e *Engine) ParseInput(addressString string) (Input, error) {
	var colorString string
	var pieceString string
	var destinationX int
	var destinationY int
	var startX int
	var startY int
	var color PieceColor
	var pieceType PieceType
	var piece *Piece
	values := strings.Split(addressString, "-")

	if len(values[0]) >= 2 {
		colorString = string(values[0][0])
		pieceString = string(values[0][1])

		// TODO: remove this check we know who is playing we do not need the color of the piece specified
		//
		switch colorString {
		case "B":
			color = Black
		case "W":
			color = White
		default:
			return Input{}, fmt.Errorf("invalid address")
		}

		switch pieceString {
		case "K":
			pieceType = King
		case "Q":
			pieceType = Queen
		case "R":
			pieceType = Rook
		case "B":
			pieceType = Bishop
		case "N":
			pieceType = Knight
		case "P":
			pieceType = Pawn
		default:
			return Input{}, fmt.Errorf("invalid address")
		}

	} else {
		return Input{}, fmt.Errorf("invalid address")
	}
	// Piece from the board
	for row := range 8 {
		for col := range 8 {
			boardPiece := e.Board[row][col]

			if boardPiece != nil && boardPiece.Color == e.CurrentPlayerColor && boardPiece.Type == pieceType {
				piece = boardPiece
				startX = row
				startY = col
			}
		}
	}

	if len(values[1]) >= 2 {
		switch values[1][0] {
		case 'a':
			destinationX = 0
		case 'b':
			destinationX = 1
		case 'c':
			destinationX = 2
		case 'd':
			destinationX = 3
		case 'e':
			destinationX = 4
		case 'f':
			destinationX = 5
		case 'g':
			destinationX = 6
		case 'h':
			destinationX = 7
		default:
			return Input{}, fmt.Errorf("invalid address")
		}
	}

	c := values[1][1]
	if c < '0' || c > '9' {
		return Input{}, fmt.Errorf("invalid coordinate")
	}
	destinationY = int(c - '0') // Ascii subtraction gives the actual number rep

	return Input{
		Color:        color,
		Type:         pieceType,
		DestinationX: destinationX,
		DestinationY: destinationY,
		StartX:       startX,
		StartY:       startY,
		Piece:        piece,
	}, nil
}
