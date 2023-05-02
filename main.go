package main

import (
	"flag"
	"log"
	"os"
	"sync"

	"golang.org/x/term"
)

type row []bool

func newRow(width int) row {
	newRow := make([]bool, width, width)
	return newRow
}

func (r *row) printRow() {
	for _, i := range *r {
		if i {
			print("#")
		} else {
			print(" ")
		}
	}
}

func (r *row) ruleResult(rule uint8, cells [3]bool) bool {
	n := 0
	for i := 2; i >= 0; i-- {
		n = n << 1
		if cells[i] {
			n += 1
		}
	}

	return (rule>>n)&0b1 == 1
}

func (r *row) updateRow(rule uint8) row {
	width := len(*r)
	update := newRow(width)

	update[0] = r.ruleResult(rule, [3]bool{(*r)[width-1], (*r)[0], (*r)[1]})
	update[width-1] = r.ruleResult(rule, [3]bool{(*r)[width-2], (*r)[width-1], (*r)[0]})

	var wg sync.WaitGroup
	for i := 1; i < width-1; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			update[i] = r.ruleResult(rule, [3]bool((*r)[i-1:i+2]))
		}()
	}
	wg.Wait()

	return update
}

func (r *row) display(rule uint8) {
	width := len(*r)
	r.printRow()
	for i := 0; i < (width / 2); i++ {
		*r = r.updateRow(rule)
		r.printRow()
	}
}

func main() {
	width, _, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatalln(err)
	}

	rule := flag.Uint("rule", 90, "cellular automaton's rule(0~255)")
	flag.Parse()

	line := newRow(width)
	line[width/2] = true

	line.display(uint8(*rule))
}
