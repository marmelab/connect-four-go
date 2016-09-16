package main

import (
	"connectfour"
	"connectfour/board"
	"connectfour/renderer"
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	game := connectfour.New()
	var column int

	fmt.Println("Connect four game started")
	renderGame(game)

	for !game.IsFinished {
		if game.CurrentPlayer == connectfour.HumanPlayer {
			fmt.Println("\x1b[91;1m● \x1b[0mYour turn to play")

			column = readColumn()

			_, err := game.HumanPlay(column - 1)
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else {
			fmt.Println("\x1b[38;5;226m● \x1b[0mComputer turn to play")
			column, err := game.ComputerPlay(connectfour.AIPlayer)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("Computer played column %d\n", column+1)
		}

		renderGame(game)

		fmt.Println("-------------------------------")

		if game.IsFinished {
			if game.Winner == connectfour.HumanPlayer {
				fmt.Println("You won")
			} else if game.Winner == connectfour.AIPlayer {
				fmt.Println("Computer won")
			} else {
				fmt.Println("Draw")
			}
		}
	}
}

func readColumn() int {
	column := -1
	fmt.Println("Which column do you want to play ?")
	for !(column >= 0 && column <= board.BoardWidth) {
		_, err := fmt.Scanf("%d", &column)
		if err != nil {
			fmt.Println(err)
		} else {
			continue
		}
	}
	return column
}

func renderGame(game connectfour.Game) {
	fmt.Print(renderer.Render(game.Board, "\x1b[91;1m\x1b[48;5;67m ● \x1b[0m", "\x1b[38;5;226m\x1b[48;5;67m ● \x1b[0m", "\x1b[97;1m\x1b[48;5;67m ● \x1b[0m", " %v "))
}
