package main

import (
	"fmt"
	"strings"
)

func main() {
	// create a tic-tac-toe board
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	board[0][0] = "X"
	board[1][1] = "O"
	board[2][2] = "X"
	board[0][2] = "O"
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join((board[i]), " "))

	}

}