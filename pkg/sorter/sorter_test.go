package sorter_test

import (
	"testing"

	"github.com/shanmugh/package-sorter/pkg/sorter"
)

const (
	MaxSingleDimension = 150
	MaxVolume          = 1_000_000
	MaxMass            = 20
)

func TestSorter(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *sorter.Package
		expected string
	}{
		{
			name: "standard package",
			input: &sorter.Package{
				Width:  10,
				Height: 10,
				Length: 10,
				Mass:   10,
			},
			expected: "STANDARD",
		},
		{
			name: "special package by volume",
			input: &sorter.Package{
				Width:  100,
				Height: 100,
				Length: 100,
				Mass:   10,
			},
			expected: "SPECIAL",
		},
		{
			name: "special package by dimension",
			input: &sorter.Package{
				Width:  150,
				Height: 10,
				Length: 10,
				Mass:   10,
			},
			expected: "SPECIAL",
		},
		{
			name: "special package by mass",
			input: &sorter.Package{
				Width:  10,
				Height: 10,
				Length: 10,
				Mass:   20,
			},
			expected: "SPECIAL",
		},
		{
			name: "rejected package by volume and mass",
			input: &sorter.Package{
				Width:  100,
				Height: 100,
				Length: 100,
				Mass:   20,
			},
			expected: "REJECTED",
		},
		{
			name: "rejected package by dimension and mass",
			input: &sorter.Package{
				Width:  150,
				Height: 10,
				Length: 10,
				Mass:   20,
			},
			expected: "REJECTED",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			sorter := sorter.NewSorter(MaxSingleDimension, MaxVolume, MaxMass)
			got := sorter.Sort(tc.input.Width, tc.input.Length, tc.input.Height, tc.input.Mass)

			if got != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, got)
			}
		})
	}
}
