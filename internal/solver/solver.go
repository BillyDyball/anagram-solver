package solver

import (
	"fmt"
	"sort"
	"strings"
)

type Group struct {
	Length int
	Words  []string
}

func Find(
	dictionary []string,
	letters string,
	depth int,
	minLength int,
	maxLength int,
	include string,
) ([]Group, error) {
	if depth < 1 {
		return nil, fmt.Errorf("depth must be at least 1")
	}
	if minLength < 1 {
		return nil, fmt.Errorf("minimum word length must be at least 1")
	}
	if maxLength < 0 {
		return nil, fmt.Errorf("maximum word length must be at least 0")
	}
	if maxLength > 0 && maxLength < minLength {
		return nil, fmt.Errorf("maximum word length must be greater than or equal to minimum word length")
	}

	requiredCharacter, err := includedCharacter(include)
	if err != nil {
		return nil, err
	}

	available, err := letterCounts(letters)
	if err != nil {
		return nil, err
	}

	wordsByLength := make(map[int][]string)
	for _, word := range dictionary {
		length := len(word)
		if length < minLength || (maxLength > 0 && length > maxLength) {
			continue
		}
		if requiredCharacter != 0 && !strings.ContainsRune(word, requiredCharacter) {
			continue
		}
		if canBuild(word, available, depth) {
			wordsByLength[length] = append(wordsByLength[length], word)
		}
	}

	lengths := make([]int, 0, len(wordsByLength))
	for length := range wordsByLength {
		lengths = append(lengths, length)
	}
	sort.Ints(lengths)

	groups := make([]Group, 0, len(lengths))
	for _, length := range lengths {
		words := wordsByLength[length]
		sort.Strings(words)
		groups = append(groups, Group{Length: length, Words: words})
	}

	return groups, nil
}

func includedCharacter(input string) (rune, error) {
	if input == "" {
		return 0, nil
	}

	characters := []rune(strings.ToLower(input))
	if len(characters) != 1 || characters[0] < 'a' || characters[0] > 'z' {
		return 0, fmt.Errorf("include must be a single A-Z character")
	}
	return characters[0], nil
}

func letterCounts(input string) ([26]int64, error) {
	var counts [26]int64
	input = strings.ToLower(input)
	if input == "" {
		return counts, fmt.Errorf("letters must not be empty")
	}

	for _, character := range input {
		if character < 'a' || character > 'z' {
			return counts, fmt.Errorf("letters must contain only A-Z characters")
		}
		counts[character-'a']++
	}

	return counts, nil
}

func canBuild(word string, available [26]int64, depth int) bool {
	var required [26]int64
	for _, character := range word {
		if character < 'a' || character > 'z' {
			return false
		}
		required[character-'a']++
	}

	multiplier := int64(depth)
	for index, count := range required {
		if count > available[index]*multiplier {
			return false
		}
	}

	return true
}
