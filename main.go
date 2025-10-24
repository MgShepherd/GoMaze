package main

func main() {
	m, err := generateMaze(10, 20)
	if err != nil {
		return
	}
	m.display()
}
