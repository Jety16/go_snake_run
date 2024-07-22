package main

import "math/rand"

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

func (g *Graph) addEdge(from int, to int, weight int) {
	g.matrix[from][to] = weight
	g.matrix[to][from] = weight
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
	graph := newGraph(size)
	for i := range size {
		for j := range size {
			graph.addEdge(i, j, (rand.Int() % 10))

		}
	}
	return &Map{
		name:   name,
		size:   size,
		web:    *graph,
		player: player,
		snake:  *newSnake(),
	}
}
