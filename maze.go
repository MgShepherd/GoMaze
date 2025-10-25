package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const verticalEdge = '|'
const horizontalEdge = '-'

type position struct {
	x, y int
}

type edge struct {
	pos1, pos2 position
}

type maze struct {
	positions [][]rune
	edges     map[edge]struct{}
}

func generateMaze(width, height int) (*maze, error) {
	if width <= 0 || height <= 0 {
		return nil, errors.New("Unable to create Maze with specified dimensions")
	}

	return &maze{positions: initPositions(width, height), edges: initEdges(width, height)}, nil
}

func (m *maze) String() string {
	var builder strings.Builder
	builder.WriteString(strings.Repeat(string(horizontalEdge), len(m.positions[0])*2+1) + "\n")
	for y := range m.positions {
		m.writeLineWithVerticalEdges(&builder, y)
		if y < len(m.positions)-1 {
			m.writeLineWithHorizontalEdges(&builder, y)
		}
	}
	builder.WriteString(strings.Repeat(string(horizontalEdge), len(m.positions[0])*2+1) + "\n")
	return builder.String()
}

func (m *maze) writeLineWithVerticalEdges(builder *strings.Builder, y int) {
	builder.WriteRune(verticalEdge)
	for x := range m.positions[y] {
		builder.WriteRune(m.positions[y][x])
		if x < len(m.positions[y])-1 {
			pos1 := position{x, y}
			pos2 := position{x + 1, y}
			_, edgeExists := m.edges[edge{pos1, pos2}]
			if edgeExists {
				builder.WriteRune(verticalEdge)
			} else {
				builder.WriteRune(' ')
			}
		}
	}

	builder.WriteString(fmt.Sprintf("%c\n", verticalEdge))
}

func (m *maze) writeLineWithHorizontalEdges(builder *strings.Builder, y int) {
	builder.WriteRune(verticalEdge)
	for x := range m.positions[y] {
		pos1 := position{x, y}
		pos2 := position{x, y + 1}
		_, edgeExists := m.edges[edge{pos1, pos2}]
		if edgeExists && x < len(m.positions[y])-1 {
			builder.WriteString(fmt.Sprintf("%c%c", horizontalEdge, horizontalEdge))
		} else if edgeExists {
			builder.WriteRune(horizontalEdge)
		} else {
			builder.WriteString("  ")
		}
	}
	builder.WriteString(fmt.Sprintf("%c\n", verticalEdge))

}

func initPositions(width, height int) [][]rune {
	positions := make([][]rune, height)
	for i := range positions {
		positions[i] = make([]rune, width)
		for j := range positions[i] {
			positions[i][j] = '.'
		}
	}
	return positions
}

func initEdges(width, height int) map[edge]struct{} {
	edges := make(map[edge]struct{})
	for y := range height {
		for x := range width {
			if x < width-1 {
				pos1 := position{x, y}
				pos2 := position{x: x + 1, y: y}
				edges[edge{pos1, pos2}] = struct{}{}
			}
			if y < height-1 {
				pos1 := position{x, y}
				pos2 := position{x: x, y: y + 1}
				edges[edge{pos1, pos2}] = struct{}{}
			}
		}
	}
	return edges
}

func clearScreen() error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
