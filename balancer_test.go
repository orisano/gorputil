package gorputil

import "testing"

func TestSequential_Balance(t *testing.T) {
	tests := []struct {
		n        int
		call     int
		expected []int
	}{
		{
			n:        10,
			call:     2,
			expected: []int{0, 1},
		},
		{
			n:        5,
			call:     7,
			expected: []int{0, 1, 2, 3, 4, 0, 1},
		},
		{
			n:        1,
			call:     5,
			expected: []int{0, 0, 0, 0, 0},
		},
	}

	for _, test := range tests {
		seq := &Sequential{}
		for i := 0; i < test.call; i++ {
			if got := seq.Balance(test.n); got != test.expected[i] {
				t.Errorf("unexpected id. expected: %v, but got: %v", test.expected[i], got)
			}
		}
	}
}
