package connectfour

import (
	"bytes"
	"strconv"
	"strings"
)

type Board [6][7]uint8

func BoardFromString(input string) Board {
const FirstPlayerChar string = "x"
const SecondPlayerChar string = "o"
	lines := strings.Split(input, "\n")
	board := Board{}
	for lineIndex, line := range lines {
		for charIndex, char := range line {
			switch string(char) {
			case FirstPlayerChar:
				board[lineIndex][charIndex] = 1
			case SecondPlayerChar:
				board[lineIndex][charIndex] = 2
			}
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
