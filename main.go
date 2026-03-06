// Package main - application entry point
package main

import (
	"github.com/DevitoDbug/chess_v1/engine"
)

func main() {
	game := engine.NewEngine()
	game.Run()
}
