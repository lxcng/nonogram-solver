package main

import (
	"log"
	"nonogram-solver/puzzle"
	"os"
)

func main() {
	args := os.Args
	for i := 1; i < len(args); i++ {
		p, err := puzzle.NewPuzzleFromConfig(args[i])
		if err != nil {
			log.Println(err)
			continue
		}
		p.Solve()
	}
}
