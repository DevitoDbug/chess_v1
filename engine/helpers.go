package engine

import (
	"fmt"
	"strconv"
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
	StartX       int32
	StartY       int32
	DestinationX int32
	DestinationY int32
}

// Expect input like a2f4
// Get the starting square and the destination square
func ParseInput(addressInputString string) (Input, error) {
	chars := strings.Split(addressInputString, "")
	fmt.Printf("Input is: %+v", chars)

	if len(chars) != 4 {
		return Input{}, fmt.Errorf("invalid input. String provided is not of the required length")
	}

	startingX, err := ConvertLetterToCorrespondingIndexNumber(chars[0])
	if err != nil {
		return Input{}, fmt.Errorf("invalid input. First char is invalid try using a letter between a and h")
	}

	startingY, err := ConvertStringNumberToInt32(chars[1])
	if err != nil {
		return Input{}, fmt.Errorf("invalid input. Second char is invalid try using a  between 1 and 8")
	}

	destinationX, err := ConvertLetterToCorrespondingIndexNumber(chars[2])
	if err != nil {
		return Input{}, fmt.Errorf("invalid input. Third char is invalid try using a letter between a and h")
	}

	destinationY, err := ConvertStringNumberToInt32(chars[3])
	if err != nil {
		return Input{}, fmt.Errorf("invalid input. Forth char is invalid try using a  between 1 and 8")
	}

	return Input{
		StartX:       startingX,
		StartY:       startingY - 1,
		DestinationX: destinationX,
		DestinationY: destinationY - 1,
	}, nil
}

func ConvertLetterToCorrespondingIndexNumber(char string) (int32, error) {
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

func ConvertStringNumberToInt32(char string) (int32, error) {
	num, err := strconv.Atoi(char)
	if err != nil || num <= 0 { // No position x = 0 in the board from the users view
		return 0, err
	}
	return int32(num), nil
}
