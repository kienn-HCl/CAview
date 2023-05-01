package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/term"
)

func printRow(row []bool, width int) {
	for i := 0; i < width; i++ {
		if row[i] {
			print(" ")
		} else {
			print("#")
		}
	}
}

func ruleResult(cells int8, rule int8) bool {
    n := 0
    for i:=0; i<3; i++ {
    }
}

func main() {
	width, height, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(width, height)

	row := make([]bool, width, width)
	row[width/2] = true
}
