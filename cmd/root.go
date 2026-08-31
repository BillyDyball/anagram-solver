package cmd

import (
	"fmt"
	"io"

	"github.com/BillyDyball/anagram-solver/internal/solver"
	"github.com/spf13/cobra"
)

func NewRootCommand(words []string, out, errOut io.Writer) *cobra.Command {
	var depth int
	var minLength int
	var maxLength int
	var include string

	command := &cobra.Command{
		Use:           "anagramsolve LETTERS",
		Short:         "Find words that can be made from a set of letters",
		Args:          cobra.ExactArgs(1),
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if depth < 1 {
				return fmt.Errorf("depth must be at least 1")
			}

			results, err := solver.Find(
				words,
				args[0],
				depth,
				minLength,
				maxLength,
				include,
			)
			if err != nil {
				return err
			}

			printResults(cmd.OutOrStdout(), results)
			return nil
		},
	}

	command.SetOut(out)
	command.SetErr(errOut)
	command.Flags().IntVarP(
		&depth,
		"depth",
		"d",
		1,
		"number of times each supplied letter may be used",
	)
	command.Flags().IntVar(
		&minLength,
		"min",
		3,
		"minimum word length",
	)
	command.Flags().IntVar(
		&maxLength,
		"max",
		0,
		"maximum word length (0 means no limit)",
	)
	command.Flags().StringVar(
		&include,
		"include",
		"",
		"character that every result must contain",
	)

	return command
}

func printResults(out io.Writer, groups []solver.Group) {
	if len(groups) == 0 {
		fmt.Fprintln(out, "No anagrams found.")
		return
	}

	for index, group := range groups {
		if index > 0 {
			fmt.Fprintln(out)
		}

		label := "letters"
		if group.Length == 1 {
			label = "letter"
		}
		fmt.Fprintf(out, "%d %s:\n", group.Length, label)
		for _, word := range group.Words {
			fmt.Fprintln(out, word)
		}
	}
}
