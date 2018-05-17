package roman

import (
	"fmt"
	"strings"
	"testing"
)

type testCase struct {
	in   string
	want int
}
type errCase struct {
	in   string
	want string
}

func TestSimpleChars(t *testing.T) {
	tcs := []testCase{
		{"", 0},
		{"I", 1},
		{"V", 5},
		{"X", 10},
		{"L", 50},
		{"C", 100},
		{"D", 500},
		{"M", 1000},
	}

	for _, tc := range tcs {
		got, err := From(tc.in)
		if err != nil {
			t.Errorf("TestSimpleChars(%q) -> %v", tc.in, err)
		}
		if got != tc.want {
			t.Errorf("TestSimpleChars(%q)=%d, want %d", tc.in, got, tc.want)
		}
	}
}
func TestSimpleChords(t *testing.T) {
	tcs := []testCase{
		{"III", 3},
		{"XVI", 16},
		{"XXXVIII", 38},
		{"LI", 51},
		{"CCCXV", 315},
		{"DV", 505},
		{"MM", 2000},
		{"MDCCCLXXXVIII", 1888},
	}

	for _, tc := range tcs {
		got, err := From(tc.in)
		if err != nil {
			t.Errorf("TestSimpleChords(%q) -> %v", tc.in, err)
		}
		if got != tc.want {
			t.Errorf("TestSimpleChords(%q)=%d, want %d", tc.in, got, tc.want)
		}
	}
}

func TestCharacterErrors(t *testing.T) {
	format := "invalid character %q in %q"

	tcs := []errCase{
		{"MMMMMm", fmt.Sprintf(format, 'm', "MMMMMm")},
		{".", fmt.Sprintf(format, '.', ".")},
	}

	for _, tc := range tcs {
		_, err := From(tc.in)

		if err.Error() != tc.want {
			t.Errorf("TestErrors(%q) -> %v, want %v", tc.in, err, tc.want)
		}
	}
}

func TestSequenceErrors(t *testing.T) {

	tcs := []errCase{
		{"IIII", "maximum occurences exceeded"},
	}

	for _, tc := range tcs {
		_, err := From(tc.in)

		if !strings.Contains(err.Error(), tc.want) {
			t.Errorf("TestErrors(%q) -> %v, want %v", tc.in, err, tc.want)
		}
	}
}
