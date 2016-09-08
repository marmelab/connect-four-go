package renderer

import (
	"bytes"
	"connectfour"
)

func Render(board connectfour.Board, firstPlayerChar, secondPlayerChar string) string {
	var buffer bytes.Buffer
	for _, line := range board {
		for _, cell := range line {
			switch cell {
			case 1:
				buffer.WriteString(firstPlayerChar)
			case 2:
				buffer.WriteString(secondPlayerChar)
			default:
				buffer.WriteString(" ")
			}
		}
		buffer.WriteString("\n")
	}
	output := buffer.String()
	return output[:len(output)-1]
}
