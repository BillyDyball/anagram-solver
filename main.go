package main

import (
	"fmt"
	"os"

	"github.com/BillyDyball/anagram-solver/cmd"
	"github.com/BillyDyball/anagram-solver/internal/dictionary"
)

func main() {
	words, err := dictionary.Words()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load embedded dictionary: %v\n", err)
		os.Exit(1)
	}

	if err := cmd.NewRootCommand(words, os.Stdout, os.Stderr).Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
