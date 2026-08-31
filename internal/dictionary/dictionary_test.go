package dictionary

import (
	"sort"
	"testing"
)

func TestEmbeddedDictionaryLoadsWordList(t *testing.T) {
	loaded, err := Words()
	if err != nil {
		t.Fatalf("Words() error = %v", err)
	}
	if len(loaded) < 300000 {
		t.Fatalf("Words() loaded %d words, want at least 300000", len(loaded))
	}
	if !sort.StringsAreSorted(loaded) {
		t.Fatal("Words() result is not sorted")
	}
	index := sort.SearchStrings(loaded, "larboard")
	if index == len(loaded) || loaded[index] != "larboard" {
		t.Fatal("Words() does not contain larboard")
	}
}
