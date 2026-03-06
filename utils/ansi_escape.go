package utils

import "fmt"

type AnsiEscapeColor string

const (
	Reset   AnsiEscapeColor = "\033[0m"
	Red     AnsiEscapeColor = "\033[31m"
	Green   AnsiEscapeColor = "\033[32m"
	Yellow  AnsiEscapeColor = "\033[33m"
	Blue    AnsiEscapeColor = "\033[34m"
	Magenta AnsiEscapeColor = "\033[35m"
	Cyan    AnsiEscapeColor = "\033[36m"
	Gray    AnsiEscapeColor = "\033[37m"
	White   AnsiEscapeColor = "\033[97m"

	AnsiClearTerminal     = "\033[2J"
	AnsiMoveCursorTopLeft = "\033[H"
)

func ClearScreen() {
	fmt.Printf("%v", AnsiClearTerminal)
}

func MoveCursorTopLeft() {
	fmt.Printf("%v", AnsiMoveCursorTopLeft)
}
