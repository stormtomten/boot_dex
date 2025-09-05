package main

import (
	//"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Whack A Mole!",
			expected: []string{"whack", "a", "mole!"},
		},
		{
			input:    "\n\t\r",
			expected: []string{},
		},
		{
			input:    "\taa\r\n   c\r",
			expected: []string{"aa", "c"},
		},
		{
			input:    "Café Pokémon",
			expected: []string{"café", "pokémon"},
		},
		{
			input:    "Hello, World!",
			expected: []string{"hello,", "world!"},
		},
		// add more cases here
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Fatalf("length missmatch: got %d, want %d(input=%q", len(actual), len(c.expected), c.input)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("word %d missmatch: got %q, want %q (input=%q)", i, word, expectedWord, c.input)
			}
		}
	}
}
