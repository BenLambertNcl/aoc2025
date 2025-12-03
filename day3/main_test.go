package main

import "testing"

func TestFindJoltage(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"987654321111111", "98"},
		{"811111111111119", "89"},
		{"234234234234278", "78"},
		{"818181911112111", "92"},
	}

	for _, tt := range tests {
		t.Run("Find Joltages", func(t *testing.T) {
			want := tt.expected
			got := findJoltage(tt.input, 2)

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}
		})
	}
}

func TestFindLargerJoltage(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"987654321111111", "987654321111"},
		{"811111111111119", "811111111119"},
		{"234234234234278", "434234234278"},
		{"818181911112111", "888911112111"},
	}

	for _, tt := range tests {
		t.Run("Find Large Joltages", func(t *testing.T) {
			want := tt.expected
			got := findJoltage(tt.input, 12)

			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}
		})
	}
}
