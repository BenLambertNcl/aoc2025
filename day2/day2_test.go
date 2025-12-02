package main

import (
	"fmt"
	"testing"
)

func TestIsInvalidId(t *testing.T) {
	var invalidTests = []string{
		"11",
		"22",
		"99",
		"1010",
		"1188511885",
		"222222",
		"446446",
		"38593859",
	}

	for _, tt := range invalidTests {
		testname := fmt.Sprintf("Check %s is invalid", tt)
		t.Run(testname, func(t *testing.T) {
			assertIdInvalid(t, tt)
		})
	}

	var validTests = []string{
		"10",
		"1111111110",
		"38593856",
		"12345678",
		"222",
	}

	for _, tt := range validTests {
		testname := fmt.Sprintf("Check %s is valid", tt)
		t.Run(testname, func(t *testing.T) {
			assertIdValid(t, tt)
		})
	}

	var invalidTestsDay2 = []string{
		"101010",
		"222",
	}

	for _, tt := range invalidTestsDay2 {
		testname := fmt.Sprintf("Check %s is invalid day2", tt)
		t.Run(testname, func(t *testing.T) {
			assertIdInvalidDay2(t, tt)
		})
	}

}

func assertIdValid(t *testing.T, id string) {
	t.Helper()
	if isValidId(id, true) != true {
		t.Errorf("isValidId(\"%s\") is false, expected true", id)
	}
}

func assertIdInvalid(t *testing.T, id string) {
	t.Helper()
	if isValidId(id, true) == true {
		t.Errorf("isValidId(\"%s\") is true, expected false", id)
	}
}

func assertIdInvalidDay2(t *testing.T, id string) {
	t.Helper()
	if isValidId(id, false) == true {
		t.Errorf("isValidId(\"%s\") is true, expected false", id)
	}
}
