package utils

import (
	"fmt"
)

// ANSI color codes
const (
	red    = "\033[41m"
	white  = "\033[47m"
	blue   = "\033[44m"
	yellow = "\033[43m"
	reset  = "\033[0m"
)

// Global variables
var (
	h, w       int
	grid       [][]int
	wallChar   = "X"
	playerChar = ">"
	awardChar  = "@"
)

func ReadInput() string {
	fmt.Scanf("%d %d", &h, &w)

	// In case if user entered negative height or width
	if h <= 0 || w <= 0 {
		return "error"
	}

	grid = make([][]int, h)
	for i := range grid {
		// Creating buffer in grid with the width of w
		grid[i] = make([]int, w)

		row := ""
		fmt.Scanf("%s", &row)

		if len(row) != w {
			return "error"
		}

		for j, char := range row {
			if !IsValidInput(char) {
				return "error"
			}

			grid[i][j] = int(char)
		}
	}

	// Reading the optional characters
	var optionalChars string
	fmt.Scanf("%s", &optionalChars)
	if len(optionalChars) == 3 {
		wallChar = string(optionalChars[0])
		playerChar = string(optionalChars[1])
		awardChar = string(optionalChars[2])
	}

	PrintHorizontalNotation(w)
	Printmap(h, w, grid)

	return "success"
}
