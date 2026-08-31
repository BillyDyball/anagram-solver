package solver

import (
	"reflect"
	"testing"
)

func TestFindGroupsAndSortsWords(t *testing.T) {
	dictionary := []string{"cab", "baa", "ba", "abc", "a", "ab", "baca"}

	got, err := Find(dictionary, "abc", 1, 1, 0, "")
	if err != nil {
		t.Fatalf("Find() error = %v", err)
	}

	want := []Group{
		{Length: 1, Words: []string{"a"}},
		{Length: 2, Words: []string{"ab", "ba"}},
		{Length: 3, Words: []string{"abc", "cab"}},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Find() = %#v, want %#v", got, want)
	}
}

func TestFindAppliesDepthToEachSuppliedLetter(t *testing.T) {
	dictionary := []string{"baa", "baca", "abacus"}

	got, err := Find(dictionary, "abc", 2, 1, 0, "")
	if err != nil {
		t.Fatalf("Find() error = %v", err)
	}

	want := []Group{
		{Length: 3, Words: []string{"baa"}},
		{Length: 4, Words: []string{"baca"}},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Find() = %#v, want %#v", got, want)
	}
}

func TestFindFiltersByWordLength(t *testing.T) {
	dictionary := []string{"cat", "acts", "tacos", "actors", "cartoons"}

	got, err := Find(dictionary, "cartoons", 1, 4, 6, "")
	if err != nil {
		t.Fatalf("Find() error = %v", err)
	}

	want := []Group{
		{Length: 4, Words: []string{"acts"}},
		{Length: 5, Words: []string{"tacos"}},
		{Length: 6, Words: []string{"actors"}},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Find() = %#v, want %#v", got, want)
	}
}

func TestFindRequiresIncludedCharacter(t *testing.T) {
	dictionary := []string{"ate", "eat", "tea", "teal", "late"}

	got, err := Find(dictionary, "late", 1, 3, 0, "L")
	if err != nil {
		t.Fatalf("Find() error = %v", err)
	}

	want := []Group{
		{Length: 4, Words: []string{"late", "teal"}},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Find() = %#v, want %#v", got, want)
	}
}

func TestFindRejectsInvalidInput(t *testing.T) {
	for _, input := range []string{"", "abc!", "café"} {
		if _, err := Find(nil, input, 1, 1, 0, ""); err == nil {
			t.Errorf("Find(nil, %q, 1, 1, 0, \"\") expected an error", input)
		}
	}
}

func TestFindRejectsInvalidIncludedCharacter(t *testing.T) {
	for _, include := range []string{"ll", "!", "é"} {
		if _, err := Find(nil, "abc", 1, 1, 0, include); err == nil {
			t.Errorf("Find() with include %q expected an error", include)
		}
	}
}

func TestFindRejectsInvalidLengthRange(t *testing.T) {
	tests := []struct {
		min int
		max int
	}{
		{min: 0, max: 0},
		{min: 3, max: -1},
		{min: 6, max: 4},
	}

	for _, test := range tests {
		if _, err := Find(nil, "abc", 1, test.min, test.max, ""); err == nil {
			t.Errorf("Find() with min %d and max %d expected an error", test.min, test.max)
		}
	}
}
