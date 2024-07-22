package main

type Graph struct {
	n      int
	matrix [][]int
}

type Map struct {
	name   string
	size   int
	web    Graph
	player Player
	snake  Snake
}

func newGraph(n int) *Graph {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
		for j := range matrix[i] {
			matrix[i][j] = 0
		}
	}
	return &Graph{
		n:      n,
		matrix: matrix,
	}
}

func newMap(name string, player Player, size int) *Map {
	return &Map{
		name:   name,
		size:   size,
		web:    *newGraph(size),
		player: player,
		snake:  *newSnake(),
	}
}
