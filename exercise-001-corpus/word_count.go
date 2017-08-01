package main

import (
	"fmt"
	"github.com/tbuist/saigo/exercise-001-corpus/corpus"
	"os"
)

func main() {
	var arg string
	if len(os.Args) != 2 {
		fmt.Println("Improper arguments given")
		os.Exit(0)
	} else {
		arg = os.Args[1]
	}

	for i, pair := range corpus.WordCount(arg) {
		fmt.Printf("%d %s %d\n", i, pair.Word, pair.Count)
	}
}
