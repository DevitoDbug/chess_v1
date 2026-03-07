// Package main - application entry point
package main

import (
	"fmt"
	"strings"
)

func main() {
	values := strings.Split("B-a2", "-")
	if len(values[0]) >= 2 {
		colorType := values[0][0]
		pieceType := values[0][1]
		fmt.Printf("color: %c \n piece: %c\n", colorType, pieceType)
	} else {
		fmt.Print("error")
	}
	// game := engine.NewEngine()
	// game.Run()
}
