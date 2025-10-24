package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

type maze struct {
	grid [][]int
}

func generateMaze(width, height int) (*maze, error) {
	if width <= 0 || height <= 0 {
		return nil, errors.New("Unable to create Maze with specified dimensions")
	}
	grid := make([][]int, height)
	for i := range grid {
		grid[i] = make([]int, width)
	}
	return &maze{grid}, nil
}

func (m *maze) display() {
	if clearScreen() != nil {
		fmt.Fprintf(os.Stderr, "There was a problem clearing the screen\n")
	}
	fmt.Printf("Width: %d, Height: %d\n", len(m.grid[0]), len(m.grid))
}

func clearScreen() error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
