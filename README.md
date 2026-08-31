# Anagram Solver

`anagramsolve` is a terminal application that finds English words which can
be built from a supplied pool of letters. Results are grouped by word length
and sorted alphabetically.

The English dictionary is embedded in the binary, so the compiled application
is a single executable with no runtime data files.

## Requirements

- Go 1.23 or newer

## Build

```sh
go build -o anagramsolve .
```

## Usage

Use each supplied letter once:

```sh
./anagramsolve abcdefg
```

Allow each supplied letter to be used up to ten times:

```sh
./anagramsolve abcdefg -d 10
```

Only include words between four and six letters long:

```sh
./anagramsolve abcdefg --min 4 --max 6
```

Require every result to contain a specific character:

```sh
./anagramsolve abcdefgl --include l
```

The minimum word length defaults to 3. The maximum has no limit unless
`--max` is provided. `--include` accepts one letter and is case-insensitive.
All filters can be combined with `--depth`.

Repeated input letters add to the available pool. For example, `letter`
provides one `l`, two `e` characters, two `t` characters, and one `r`.
The depth multiplier is then applied to those counts.

Input is case-insensitive and accepts the letters A-Z.

## Test

```sh
go test ./...
```

## Dictionary

The lowercase English word list in
`internal/dictionary/data/words_alpha.txt` is embedded directly into the
executable.
