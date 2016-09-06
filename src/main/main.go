package main

import (
	"connectfour"
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
	dat, err := ioutil.ReadFile(absoluteFilePath)
	check(err)

	board := connectfour.BoardFromString(string(dat))

	fmt.Println(connectfour.BoardToString(board))
}
