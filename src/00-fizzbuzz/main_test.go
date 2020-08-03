package main

import "testing"

var tests = []struct {
	count    int
	f        int
	b        int
	expected string
}{
	{1, 3, 5, 1},
}

func TestFizzBuzz(t *testing.T) {
	for _, tt := range tests {
		res := fizzBuzz(tt.count, tt.f, tt.b)
		if res != tt.expected {
			t.Errorf("expected %v to be printed for count: %v. got %v", res, tt.count, tt.expected)
		}
	}
}
