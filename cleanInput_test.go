package main

import "testing"

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
			input:    "PIKACHU",
			expected: []string{"pikachu"},
		},
		{
			input:    "  Bulbasaur   Charmander  ",
			expected: []string{"bulbasaur", "charmander"},
		},
	}

	for _, c := range cases {

		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf(
				"expected length %d got %d",
				len(c.expected),
				len(actual),
			)
			continue
		}

		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf(
					"expected %s got %s",
					c.expected[i],
					actual[i],
				)
			}
		}
	}
}