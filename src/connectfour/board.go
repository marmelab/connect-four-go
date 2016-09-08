package connectfour

import (
	"errors"
)

const BoardWidth int = 7
const BoardHeight int = 6

type Board [BoardHeight][BoardWidth]int

func (board Board) AddDisc(column, player int) (Board, error) {
	dropped := false

	for i := BoardHeight - 1; i >= 0; i-- {
		if board[i][column] == 0 {
			dropped = true
			board[i][column] = player
			break
		}
	}

	if !dropped {
		return board, errors.New("Illegal move")
	}

	return board, nil
}
