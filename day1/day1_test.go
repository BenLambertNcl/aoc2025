package main

import "testing"

func TestAmountRotation(t *testing.T) {
	t.Run("Gets left rotation direction", func(t *testing.T) {
		want := Direction(0)
		got, _ := getRotationAmount("L10")
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("Gets right rotation direction", func(t *testing.T) {
		want := Direction(1)
		got, _ := getRotationAmount("R10")
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("Gets rotation amount", func(t *testing.T) {
		var want int64 = 10
		_, got := getRotationAmount("R10")
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("Gets rotation amount above max", func(t *testing.T) {
		var want int64 = 1000
		_, got := getRotationAmount("R1000")
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestDay1(t *testing.T) {
	t.Run("Day1 counts zeros correctly", func(t *testing.T) {
		want := 1
		got := part1([]string{"L50"})
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("Day1 counts zeros correctly multiple rotations", func(t *testing.T) {
		want := 1
		got := part1([]string{"L150"})
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
