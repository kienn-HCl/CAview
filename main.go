package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/term"
)

func printRow(row []bool) {
	for _, i := range row {
		if i {
			print("#")
		} else {
			print(" ")
		}
	}
}

func ruleResult(cells [3]bool, rule int8) bool {
    n := 0
    for i:=2; i>=0; i-- {
        n = n << 1
        if cells[i] {
            n+=1
        } 
    }

    return (rule >> n) & 0b1 == 1
}

func updateRow(row []bool, rule int8, width int) []bool {
    update := make([]bool, width, width)

    update[0] = ruleResult([3]bool{row[width-1],row[0],row[1]}, rule)
    update[width-1] = ruleResult([3]bool{row[width-2],row[width-1],row[0]}, rule)
    for i:=1; i<width-1; i++ {
        update[i] = ruleResult([3]bool(row[i-1:i+2]), rule)
    }

    return update
}

func main() {
	width, height, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(width, height)

	row := make([]bool, width, width)
	row[width/2] = true

    printRow(row)
    for i:= 0; i<(width/2); i++ {
        row = updateRow(row, 90, width)
        printRow(row)
    }
}
