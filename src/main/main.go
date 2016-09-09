package main

import (
	"connectfour"
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
			fmt.Println("Your turn to play")
			fmt.Println("Which column do you want to play ?")
			_, err := fmt.Scanf("%d\n", &column)
			_, err = game.HumanPlay(column - 1)
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else {
			fmt.Println("Computer turn to play")
			column, err := game.ComputerPlay()
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("Computer played column %d\n", column)
		}

		renderGame(game)

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

func renderGame(game connectfour.Game) {
	renderer.Render(game.Board, "x", "o", " ")
}
