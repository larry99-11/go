package main

import (
	"testing"
)

// unit test for our functions
func TestCleanInput(t *testing.T) {

	// we are going to create a slice of 'case' structs
	cases := []struct {
		input    string
		expected []string
	}{
		// we can add test cases here
		{
			input: "hello world",
			expected: []string{
				"hello",
				"world",
			},
		},
		// upper case check
		{
			input: "HELLO World",
			expected: []string{
				"hello",
				"world",
			},
		},
	}

	// checking for the length of the arrays i.e the slices
	for _, cases := range cases {
		actual := cleanInput(cases.input)
		if len(actual) != len(cases.expected) {
			t.Errorf("lengths are not equal: %v and %v", len(actual), len(cases.expected))
			continue
		}
		// checking for the contents in our function we are testing
		for i := range actual { //actual data type is a slice
			actualWord := actual[i]
			expectedWord := cases.expected[i]
			if actualWord != expectedWord {
				t.Errorf("%v does not equal %v", actualWord, expectedWord)
			}
		}

	}
}
