package main

import (
	"reflect"
	"testing"
)

func TestShouldCreateNewMaze(t *testing.T) {
	width, height := 3, 3
	expectedMaze := &maze{grid: [][]rune{
		{'#', '#', '#'},
		{'#', '.', '#'},
		{'#', '#', '#'},
	}}

	maze, err := generateMaze(width, height)
	if err != nil {
		t.Errorf("Unexpected error when creating maze: %s\n", err)
	}
	if !reflect.DeepEqual(expectedMaze, maze) {
		t.Errorf("Expected maze \n%s but got \n%s\n", expectedMaze, maze)
	}
}

func TestReturnErrorWhenCreatingMazeWithInvalidSize(t *testing.T) {
	var errorCases = []struct {
		name          string
		width, height int
	}{
		{name: "Negative Width", width: -1, height: 10},
		{name: "Negative Height", width: 10, height: -1},
		{name: "Zero width", width: 0, height: 10},
		{name: "Zero height", width: 10, height: -1},
	}

	for _, c := range errorCases {
		t.Run(c.name, func(t *testing.T) {
			_, err := generateMaze(c.width, c.height)
			if err == nil {
				t.Errorf("Expected error when creating maze with invalid size, but none found.")
			}
		})
	}
}
