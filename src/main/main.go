package main

import (
	"connectfour/parser"
	"connectfour/renderer"
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
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

	fmt.Println(renderer.Render(board, "x", "o", " "))
}
