package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestCommandPrintsGroupedResults(t *testing.T) {
	var output bytes.Buffer
	command := NewRootCommand(
		[]string{"cab", "baa", "ba", "abc", "a", "ab"},
		&output,
		&output,
	)
	command.SetArgs([]string{"abc"})

	if err := command.Execute(); err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	want := "3 letters:\nabc\ncab\n"
	if output.String() != want {
		t.Fatalf("output = %q, want %q", output.String(), want)
	}
}

func TestCommandFiltersByMinimumAndMaximumLength(t *testing.T) {
	var output bytes.Buffer
	command := NewRootCommand(
		[]string{"cat", "acts", "tacos", "actors", "cartoons"},
		&output,
		&output,
	)
	command.SetArgs([]string{"cartoons", "--min", "4", "--max", "6"})

	if err := command.Execute(); err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	want := "4 letters:\nacts\n\n5 letters:\ntacos\n\n6 letters:\nactors\n"
	if output.String() != want {
		t.Fatalf("output = %q, want %q", output.String(), want)
	}
}

func TestCommandSupportsDepthShorthand(t *testing.T) {
	var output bytes.Buffer
	command := NewRootCommand([]string{"baa"}, &output, &output)
	command.SetArgs([]string{"abc", "-d", "2"})

	if err := command.Execute(); err != nil {
		t.Fatalf("Execute() error = %v", err)
	}
	if !strings.Contains(output.String(), "baa") {
		t.Fatalf("output %q does not contain baa", output.String())
	}
}

func TestCommandFiltersByIncludedCharacter(t *testing.T) {
	var output bytes.Buffer
	command := NewRootCommand(
		[]string{"ate", "eat", "tea", "teal", "late"},
		&output,
		&output,
	)
	command.SetArgs([]string{"late", "--include", "l"})

	if err := command.Execute(); err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	want := "4 letters:\nlate\nteal\n"
	if output.String() != want {
		t.Fatalf("output = %q, want %q", output.String(), want)
	}
}

func TestCommandRejectsInvalidIncludedCharacter(t *testing.T) {
	command := NewRootCommand(nil, &bytes.Buffer{}, &bytes.Buffer{})
	command.SetArgs([]string{"abcdef", "--include", "ab"})

	if err := command.Execute(); err == nil {
		t.Fatal("Execute() expected an error")
	}
}

func TestCommandRejectsMaximumBelowMinimum(t *testing.T) {
	command := NewRootCommand(nil, &bytes.Buffer{}, &bytes.Buffer{})
	command.SetArgs([]string{"abcdef", "--min", "6", "--max", "4"})

	if err := command.Execute(); err == nil {
		t.Fatal("Execute() expected an error")
	}
}

func TestCommandRejectsNonPositiveDepth(t *testing.T) {
	command := NewRootCommand(nil, &bytes.Buffer{}, &bytes.Buffer{})
	command.SetArgs([]string{"abc", "-d", "0"})

	if err := command.Execute(); err == nil {
		t.Fatal("Execute() expected an error")
	}
}
