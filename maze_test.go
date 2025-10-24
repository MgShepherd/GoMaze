package main

import "testing"

func TestShouldCreateNewMazeOfSize(t *testing.T) {
	width, height := 10, 20
	maze, err := generateMaze(width, height)
	if err != nil {
		t.Errorf("Unexpected error when creating maze: %s\n", err)
	}
	if len(maze.grid) != height || len(maze.grid[0]) != width {
		t.Errorf("Maze created with unexpected size, expected (%d, %d), got (%d, %d)", width, height, len(maze.grid[0]), len(maze.grid))
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
