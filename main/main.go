package main

import (
	"checkers/checkers/checkers"
	"fmt"
)

func main() {

	game := checkers.New()

	turn := game.Turn

	color := turn.Color

	fmt.Println(color)

}
