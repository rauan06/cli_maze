package main

import (
	"fmt"

	"github.com/alem-platform/ap"
)

// ANSI color codes
const (
	red    = "\033[41m"
	white  = "\033[47m"
	blue   = "\033[44m"
	yellow = "\033[43m"
	reset  = "\033[0m"
)

var (
	h, w       int
	grid       [][]int
	wallChar   = "X"
	playerChar = ">"
	awardChar  = "@"
)

func main() {
	readInput()
	printHorizontalNotation(w)
	printmap(h, w, grid)
	ap.PutRune('\n')
}

func readInput() {
	fmt.Scanf("%d %d", &h, &w)

	// Reading the main
	grid = make([][]int, h)
	for i := range grid {
		grid[i] = make([]int, w)
		var row string
		fmt.Scanf("%s", &row)
		for j, char := range row {
			// Convert ASCII digit to integer
			grid[i][j] = int(char - '0')
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
}

func printHorizontalNotation(w int) {
	blank := "      "
	ap.PutRune('\n')
	for j := 1; j <= w; j++ {
		temp := j / 26
		for _, r := range blank {
			ap.PutRune(r)
		}
		if temp == 0 {
			ap.PutRune(' ')
		} else {
			ap.PutRune(rune(temp + 64))
		}
		ap.PutRune(rune((j % 26) + 64))
	}
	ap.PutRune('\n')

	// Top of the Graph
	ap.PutRune(' ')
	ap.PutRune(' ')
	ap.PutRune(' ')
	ap.PutRune(' ')
	for j := 0; j < w*7+w-1; j++ {
		ap.PutRune('_')
	}
	ap.PutRune('\n')
}

func printmap(h, w int, grid [][]int) {
	// Walls of the Graph
	oneWallPart := "       "
	oneWallBottom := "_______"
	zeroWall := wallChar + wallChar + wallChar + wallChar + wallChar + wallChar + wallChar
	twoWall := "   " + playerChar + "   "
	threeWall := "   " + awardChar + "   "

	// Drawing the walls of the Graph
	for i := 0; i < h; i++ {
		for c := 0; c < 3; c++ {
			printVerticalNotation(c, i+1)
			for j := 0; j < w; j++ {
				if j == 0 {
					ap.PutRune('|')
				}
				switch grid[i][j] {
				case 0:
					printColoredString(red, zeroWall)
				case 2:
					if c == 1 {
						printColoredString(blue, twoWall)
					} else if c == 0 {
						printColoredString(blue, oneWallPart)
					} else if c == 2 {
						printColoredString(blue, oneWallBottom)
					}
				case 3:
					if c == 1 {
						printColoredString(yellow, threeWall)
					} else if c == 0 {
						printColoredString(yellow, oneWallPart)
					} else if c == 2 {
						printColoredString(yellow, oneWallBottom)
					}
				default:
					if c != 2 {
						printColoredString(white, oneWallPart)
					} else {
						printColoredString(white, oneWallBottom)
					}
				}
				ap.PutRune('|')
			}
			ap.PutRune('\n')
		}
	}
}

func printColoredString(color, str string) {
	for _, r := range color + str + reset {
		ap.PutRune(r)
	}
}

func printVerticalNotation(row, number int) {
	if row != 1 {
		ap.PutRune(' ')
		ap.PutRune(' ')
		ap.PutRune(' ')
	} else {
		if number < 10 {
			ap.PutRune(' ')
			PrintNumber(number)
			ap.PutRune(' ')
		} else if number >= 10 && number < 100 {
			PrintNumber(number)
			ap.PutRune(' ')
		} else if number >= 100 && number < 1000 {
			PrintNumber(number)
		}
	}
}

func PrintNumber(num int) {
	digits := []rune("0123456789")
	if num < 0 {
		ap.PutRune('-')
		num *= -1
	}

	if num < 10 {
		ap.PutRune(digits[num])
		return
	}

	PrintNumber(num / 10)
	ap.PutRune(digits[num%10])
}