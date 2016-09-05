package main

import (
	"fmt"
    "flag"
    "strings"
    "strconv"
    "io/ioutil"
    "path/filepath"
)

type Board [6][7]uint8

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
    board := stringToBoard(string(dat))

    printBoard(board)
}

func stringToBoard(input string) Board {
    lines := strings.Split(input, "\n")
    board := Board{}
    for lineIndex,line := range lines {
        for charIndex, char := range line {
            valeur, _ := strconv.Atoi(string(char))
            board[lineIndex][charIndex] = uint8(valeur)
        }
    }
    return board
}

func printBoard(board Board){
    for _, line := range board {
        for _, cell := range line {
            fmt.Printf("%v", cell)
        }
        fmt.Println()
    }
}
