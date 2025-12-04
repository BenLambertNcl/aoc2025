package main

import "testing"

func TestAdjacentPositions(t *testing.T) {
	grid := [][]string{
		{"@", "@", "@"},
		{"@", ".", "@"},
		{"@", "@", "@"},
	}

	want := 8
	got := checkAdjacentPositions(grid, 1, 1)

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestAdjacentPositionsSomeMissing(t *testing.T) {
	grid := [][]string{
		{"@", "@", "."},
		{".", ".", "@"},
		{"@", ".", "@"},
	}

	want := 5
	got := checkAdjacentPositions(grid, 1, 1)

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestAdjacentPositionsAllMissing(t *testing.T) {
	grid := [][]string{
		{".", ".", "."},
		{".", ".", "."},
		{".", ".", "."},
	}

	want := 0
	got := checkAdjacentPositions(grid, 1, 1)

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestAdjacentPositionsDoesntGoOutOfBounds(t *testing.T) {
	grid := [][]string{
		{".", ".", "."},
		{".", "@", "."},
		{".", ".", "."},
	}

	want := 1
	got := checkAdjacentPositions(grid, 0, 0)

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestPart1(t *testing.T) {
	grid := [][]string{
		{"@", "@", "@"},
		{"@", "@", "@"},
		{"@", "@", "@"},
	}

	want := 4
	got := part1(grid)

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
