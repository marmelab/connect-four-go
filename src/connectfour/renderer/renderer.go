package renderer

import (
	"bytes"
	"connectfour"
)

func ToString(board connectfour.Board) string {
	var buffer bytes.Buffer
	for _, line := range board {
		for _, cell := range line {
			switch cell {
			case 1:
				buffer.WriteString(connectfour.FirstPlayerChar)
			case 2:
				buffer.WriteString(connectfour.SecondPlayerChar)
			default:
				buffer.WriteString(" ")
			}
		}
		buffer.WriteString("\n")
	}
	output := buffer.String()
	return output[:len(output)-1]
}
