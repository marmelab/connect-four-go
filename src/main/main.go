package main

import (
	"connectfour/ai"
	"connectfour/parser"
	"connectfour/renderer"
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var filePath string
	flag.StringVar(&filePath, "file", "default", "Game filepath")
	flag.Parse()

	absoluteFilePath, _ := filepath.Abs(filePath)
	fileData, error := ioutil.ReadFile(absoluteFilePath)
	check(error)

	board := parser.Parse(string(fileData), "x", "o", " ")

	fmt.Println("Board given")
	fmt.Println(renderer.Render(board, "x", "o", " "))

	column, err := ai.NextBestMoveInTime(board, 1, 10*time.Second)
	check(err)

	board, err = board.AddDisc(column, 1)
	check(err)

	fmt.Println("Advised play : ", column)
	fmt.Println(renderer.Render(board, "x", "o", " "))
}
