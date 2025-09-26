package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HELLO WORLD",
			expected: []string{"hello", "world"},
		},
		{
			input:    "   ",
			expected: []string{},
		},
		{
			input:    "  hellO WorLD  ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Lengths don't match: '%v' vs '%v'", actual, c.expected)
			continue
		}
		for i, word := range actual {
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Words don't match: '%v' vs '%v'", word, expectedWord)
				continue
			}
		}
	}
}
