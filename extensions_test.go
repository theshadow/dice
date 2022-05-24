package dice

import (
	"fmt"
	"github.com/theshadow/dice/formula"
	"strings"
	"testing"
)

func TestDropExtension_New(t *testing.T) {

	cases := []struct {
		name     string
		which    string
		results  Results
		roll     formula.Roll
		expected string
	}{
		{
			name:    "Drop lowest without duplicates",
			which:   "lowest",
			results: Results{Rolls: []int{1, 2, 3, 4, 5}},
			roll: formula.Roll{
				Count:      5,
				Sides:      5,
				Modifier:   0,
				Extensions: make(map[string][]string),
			},
			expected: fmt.Sprintf("Dropping the lowest roll, new rolls: %v", []int{2, 3, 4, 5}),
		},
		{
			name:    "Drop lowest with duplicates",
			which:   "lowest",
			results: Results{Rolls: []int{1, 1, 3, 4, 5}},
			roll: formula.Roll{
				Count:      5,
				Sides:      5,
				Modifier:   0,
				Extensions: make(map[string][]string),
			},
			expected: fmt.Sprintf("Dropping the lowest roll, new rolls: %v", []int{1, 3, 4, 5}),
		},
		{
			name:    "Drop lowest with lowest in the middle",
			which:   "lowest",
			results: Results{Rolls: []int{2, 3, 1, 4, 5}},
			roll: formula.Roll{
				Count:      5,
				Sides:      5,
				Modifier:   0,
				Extensions: make(map[string][]string),
			},
			expected: fmt.Sprintf("Dropping the lowest roll, new rolls: %v", []int{2, 3, 4, 5}),
		},
		{
			name:    "Drop highest without duplicates",
			which:   "highest",
			results: Results{Rolls: []int{1, 2, 3, 4, 5}},
			roll: formula.Roll{
				Count:      5,
				Sides:      5,
				Modifier:   0,
				Extensions: make(map[string][]string),
			},
			expected: fmt.Sprintf("Dropping the highest roll, new rolls: %v", []int{1, 2, 3, 4}),
		},
		{
			name:    "Drop highest with duplicates",
			which:   "highest",
			results: Results{Rolls: []int{1, 2, 3, 5, 5}},
			roll: formula.Roll{
				Count:      5,
				Sides:      5,
				Modifier:   0,
				Extensions: make(map[string][]string),
			},
			expected: fmt.Sprintf("Dropping the highest roll, new rolls: %v", []int{1, 2, 3, 5}),
		},
		{
			name:    "Drop highest with highest in the middle",
			which:   "highest",
			results: Results{Rolls: []int{1, 2, 5, 3, 4}},
			roll: formula.Roll{
				Count:      5,
				Sides:      5,
				Modifier:   0,
				Extensions: make(map[string][]string),
			},
			expected: fmt.Sprintf("Dropping the highest roll, new rolls: %v", []int{1, 2, 3, 4}),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ext := newDropxtension([]string{c.which})

			actual, err := ext.Exec(c.results, c.roll)
			if err != nil {
				t.Fatalf("unexpected error while executing extension: %s", err)
			}

			if strings.Compare(c.expected, actual) != 0 {
				t.Fatalf("Extension result doesn't match, expected '%s' and got '%s'", c.expected, actual)
			}
		})
	}
}
