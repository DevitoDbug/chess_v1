package engine

import (
	"fmt"
	"strconv"
	"strings"
)

var whitePieces = [...]rune{
	Pawn:   '♟',
	Knight: '♞',
	Bishop: '♝',
	Rook:   '♜',
	Queen:  '♛',
	King:   '♚',
}

var blackPieces = [...]rune{
	Pawn:   '♙',
	Knight: '♘',
	Bishop: '♗',
	Rook:   '♖',
	Queen:  '♕',
	King:   '♔',
}

func GetRenderLetter(pieceType PieceType, color PieceColor) rune {
	switch color {
	case Black:
		return blackPieces[pieceType]
	case White:
		return whitePieces[pieceType]
	}
	return ' '
}

func getColorLetter(color PieceColor) string {
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
	StartX         int32
	StartY         int32
	DestinationX   int32
	DestinationY   int32
	PromotionPiece *PieceType
}

// Expect input like a2f4
// Get the starting square and the destination square
func parseInput(addressInputString string) (Input, error) {
	chars := strings.Split(addressInputString, "")
	var promotionPiece *PieceType

	if len(chars) != 4 {
		if len(chars) != 5 {
			return Input{}, fmt.Errorf("invalid input. String provided is not of the required length")
		}

		piece, err := getPromortionPieceType(chars[len(chars)-1])
		if err != nil {
			return Input{}, err
		}

		promotionPiece = piece
	}

	startingX, err := convertLetterToCorrespondingIndexNumber(chars[0])
	if err != nil {
		return Input{}, fmt.Errorf("invalid input. First char is invalid try using a letter between a and h")
	}

	startingY, err := convertStringNumberToInt32(chars[1])
	if err != nil {
		return Input{}, fmt.Errorf("invalid input. Second char is invalid try using a  between 1 and 8")
	}

	destinationX, err := convertLetterToCorrespondingIndexNumber(chars[2])
	if err != nil {
		return Input{}, fmt.Errorf("invalid input. Third char is invalid try using a letter between a and h")
	}

	destinationY, err := convertStringNumberToInt32(chars[3])
	if err != nil {
		return Input{}, fmt.Errorf("invalid input. Forth char is invalid try using a  between 1 and 8")
	}

	return Input{
		StartX:         startingX,
		StartY:         startingY - 1,
		DestinationX:   destinationX,
		DestinationY:   destinationY - 1,
		PromotionPiece: promotionPiece,
	}, nil
}

func convertLetterToCorrespondingIndexNumber(char string) (int32, error) {
	switch char {
	case "a":
		return 0, nil
	case "b":
		return 1, nil
	case "c":
		return 2, nil
	case "d":
		return 3, nil
	case "e":
		return 4, nil
	case "f":
		return 5, nil
	case "g":
		return 6, nil
	case "h":
		return 7, nil
	default:
		return 0, fmt.Errorf("invalid address")
	}
}

func convertStringNumberToInt32(char string) (int32, error) {
	num, err := strconv.Atoi(char)
	if err != nil || num <= 0 { // No position x = 0 in the board from the users view
		return 0, err
	}
	return int32(num), nil
}

func getPromortionPieceType(char string) (*PieceType, error) {
	switch char {
	case "q":
		piece := Queen
		return &piece, nil
	case "r":
		piece := Rook
		return &piece, nil
	case "b":
		piece := Bishop
		return &piece, nil
	case "n":
		piece := Knight
		return &piece, nil
	default:
		return nil, fmt.Errorf("promotion error, invalid piece")
	}
}

func toggleCurrentPlayer(currentPlayerColor PieceColor) PieceColor {
	if currentPlayerColor == Black {
		return White
	}

	return Black
}
