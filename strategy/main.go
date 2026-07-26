package main

import (
	"flag"
	"fmt"
)

type printer struct {
	stategy outputStrategy
}

func (rcv *printer) setStrategy(s outputStrategy) {
	rcv.stategy = s
}

func (rcv *printer) print(input string) {
	output := rcv.stategy.createOutput(input)
	fmt.Println(output)
}

func main() {
	input := flag.String("i", "", "the input")
	strat := flag.String("s", "", "the strategy")

	flag.Parse()

	p := printer{}

	switch {
	case *strat == "string":
		// ex: go run . -i "hello" -s string
		p.setStrategy(stringStrategy{})
	case *strat == "byte":
		// ex: go run . -i "hello" -s byte
		p.setStrategy(byteStrategy{})
	case *strat == "rune":
		// ex: go run . -i "hello" -s rune
		p.setStrategy(runeStrategy{})
	}

	p.print(*input)
}
