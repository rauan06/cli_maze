package utils

import "github.com/alem-platform/ap"

func PrintHorizontalNotation(w int) {
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
	for i := 0; i < 4; i++ {
		ap.PutRune(' ')
	}

	for j := 0; j < w*7+w-1; j++ {
		ap.PutRune('_')
	}
	ap.PutRune('\n')
}

func Printmap(h, w int, grid [][]int) {
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
