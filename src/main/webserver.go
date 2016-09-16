package main

import (
	"connectfour"
	"connectfour/board"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getBoardFromQueryString(query url.Values) (board.Board, error) {
	queryString := query.Get("grid")
	grid := []byte(queryString)

	var board board.Board

	err := json.Unmarshal(grid, &board)

	return board, err
}

func getAiPlayerFromQueryString(query url.Values) (int, error) {
	queryString := query.Get("aiPlayer")

	return strconv.Atoi(queryString)
}

func connectFour(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	board, err := getBoardFromQueryString(query)
	check(err)
	aiPlayer, err := getAiPlayerFromQueryString(query)
	check(err)

	game := connectfour.New()
	game.Board = board

	column, err := game.ComputerPlay(aiPlayer)
	check(err)

	if err := json.NewEncoder(w).Encode(column); err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", connectFour)
	http.ListenAndServe(":8000", nil)
}
