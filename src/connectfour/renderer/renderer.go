package renderer

import (
	"connectfour/board"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
)

func Render(gameBoard board.Board, firstPlayerChar, secondPlayerChar, emptyCellChar string) {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)
	headers := make([]string, 0)
	for i := 1; i <= board.BoardWidth; i++ {
		headers = append(headers, strconv.Itoa(i))
	}
	table.SetHeader(headers)

	for y := 0; y < board.BoardHeight; y++ {
		line := make([]string, 0)
		for x := 0; x < board.BoardWidth; x++ {
			switch gameBoard[y][x] {
			case 1:
				line = append(line, firstPlayerChar)
			case 2:
				line = append(line, secondPlayerChar)
			default:
				line = append(line, emptyCellChar)
			}
		}
		table.Append(line)
	}
	table.Render()
}
