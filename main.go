package main

import "fmt"

func main() {
	m, err := generateMaze(20, 10)
	if err != nil {
		return
	}
	fmt.Printf("%s\n", m)
}
