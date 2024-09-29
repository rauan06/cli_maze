package main

import (
	"crunch01/utils"

	"github.com/alem-platform/ap"
)

func main() {
	err := utils.ReadInput()
	if err == "error" {
		for _, r := range "Error: invalid inputs are given, try again\n" {
			ap.PutRune(r)
		}
		return
	}

	ap.PutRune('\n')
}
