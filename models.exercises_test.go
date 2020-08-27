package main

import "testing"

// Test the function that gets the exercises
func TestGetAllExercises(t *testing.T) {
	alist := getAllExercises()

	// Check that the length of the list of exercises is equal to 10
	if len(alist) != 10 {
		t.Fail()
	}
}
