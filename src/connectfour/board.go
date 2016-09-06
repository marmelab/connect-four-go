package connectfour

import (
	"bytes"
	"strconv"
	"strings"
)

type Board [6][7]uint8

func BoardFromString(input string) Board {
	lines := strings.Split(input, "\n")
	board := Board{}
	for lineIndex, line := range lines {
		for charIndex, char := range line {
			valeur, _ := strconv.Atoi(string(char))
			board[lineIndex][charIndex] = uint8(valeur)
		}
	}
	return board
}

func BoardToString(board Board) string {
	var buffer bytes.Buffer
	for _, line := range board {
		for _, cell := range line {
			buffer.WriteString(strconv.Itoa(int(cell)))
		}
		buffer.WriteString("\n")
	}
	output := buffer.String()
	return output[:len(output)-1]
}
