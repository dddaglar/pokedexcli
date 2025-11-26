package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{"", []string{}},
		{"   ", []string{}},
		{"command arg1 arg2", []string{"command", "arg1", "arg2"}},
		{"  command   arg1   arg2  ", []string{"command", "arg1", "arg2"}},
		{"singleword", []string{"singleword"}},
	}

	for _, c := range cases {
		result := cleanInput(c.input)
		if !reflect.DeepEqual(len(result), len(c.expected)) {
			t.Errorf("length of cleanInput(%q) = %v; want %v", c.input, len(result), len(c.expected))
			continue
		}
		for i := range result {
			if !reflect.DeepEqual(result[i], c.expected[i]) {
				t.Errorf("cleanInput(%q) = %v; want %v", c.input, result, c.expected)
				break
			}
		}
	}
}
