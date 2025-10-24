package main

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

type maze struct {
	grid [][]rune
}

func generateMaze(width, height int) (*maze, error) {
	if width <= 0 || height <= 0 {
		return nil, errors.New("Unable to create Maze with specified dimensions")
	}
	grid := make([][]rune, height)
	for i := range grid {
		grid[i] = make([]rune, width)
	}

	fillGrid(grid)
	return &maze{grid}, nil
}

func (m *maze) String() string {
	var builder strings.Builder
	for y := range m.grid {
		builder.WriteString(string(m.grid[y]))
		builder.WriteRune('\n')
	}
	return builder.String()
}

func fillGrid(grid [][]rune) {
	for y := range grid {
		for x := range grid[y] {
			if y == 0 || x == 0 || y == len(grid)-1 || x == len(grid[y])-1 {
				grid[y][x] = '#'
			} else {
				grid[y][x] = '.'
			}
		}
	}
}

func clearScreen() error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
