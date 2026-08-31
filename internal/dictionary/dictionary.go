package dictionary

import (
	"bufio"
	_ "embed"
	"fmt"
	"sort"
	"strings"
	"sync"
)

//go:embed data/words_alpha.txt
var dictionaryData string

var (
	loadOnce sync.Once
	words    []string
	loadErr  error
)

func Words() ([]string, error) {
	loadOnce.Do(func() {
		words, loadErr = parse(dictionaryData)
	})
	return words, loadErr
}

func parse(data string) ([]string, error) {
	found := make(map[string]struct{})
	scanner := bufio.NewScanner(strings.NewReader(data))
	scanner.Buffer(make([]byte, 64*1024), 1024*1024)

	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word == "" {
			continue
		}
		if !isLowerASCIIWord(word) {
			return nil, fmt.Errorf("dictionary contains invalid word %q", word)
		}
		found[word] = struct{}{}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	result := make([]string, 0, len(found))
	for word := range found {
		result = append(result, word)
	}
	sort.Strings(result)
	return result, nil
}

func isLowerASCIIWord(word string) bool {
	for _, character := range word {
		if character < 'a' || character > 'z' {
			return false
		}
	}
	return word != ""
}
