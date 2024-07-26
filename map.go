package main

import (
	"container/heap"
	"fmt"
	"math"
)

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

type PriorityQueueItem struct {
	x        int
	y        int
	priority int
	index    int
}

type PriorityQueue []*PriorityQueueItem

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].priority < pq[j].priority }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i]; pq[i].index, pq[j].index = i, j }

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*PriorityQueueItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (g *Graph) addEdge(fromX, fromY, toX, toY, weight int) {
	if fromX < 0 || fromX >= g.n || toX < 0 || toX >= g.n || fromY < 0 || fromY >= g.n || toY < 0 || toY >= g.n {
		return
	}
	from := fromX*g.n + fromY
	to := toX*g.n + toY
	g.matrix[from][to] = weight
	g.matrix[to][from] = weight
}

func newGraph(size int) *Graph {
	matrix := make([][]int, size*size)
	for i := range matrix {
		matrix[i] = make([]int, size*size)
		for j := range matrix[i] {
			matrix[i][j] = math.MaxInt32 // Use MaxInt32 for no edge
		}
	}
	return &Graph{
		n:      size,
		matrix: matrix,
	}
}

func newMap(name string, player Player, size int) *Map {
	graph := newGraph(size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if j < size-1 {
				graph.addEdge(i, j, i, j+1, 10)
			}
			if i < size-1 {
				graph.addEdge(i, j, i+1, j, 10)
			}
			if j > 0 {
				graph.addEdge(i, j, i, j-1, 10)
			}
			if i > 0 {
				graph.addEdge(i, j, i-1, j, 10)
			}
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

func (g *Graph) shortestPath(startX, startY, targetX, targetY int) ([][]int, int) {
	// Ensure start and target are within bounds
	if startX < 0 || startX >= g.n || targetX < 0 || targetX >= g.n || startY < 0 || startY >= g.n || targetY < 0 || targetY >= g.n {
		fmt.Println("Start or target is out of bounds")
		return nil, -1
	}

	dist := make([][]int, g.n)
	for i := range dist {
		dist[i] = make([]int, g.n)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt32
		}
	}
	dist[startX][startY] = 0

	pq := make(PriorityQueue, 1)
	pq[0] = &PriorityQueueItem{
		x:        startX,
		y:        startY,
		priority: 0,
	}
	heap.Init(&pq)

	prev := make([][]int, g.n)
	for i := range prev {
		prev[i] = make([]int, g.n)
		for j := range prev[i] {
			prev[i][j] = -1
		}
	}

	directions := [][]int{
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left
		{-1, 0}, // up
	}

	for pq.Len() > 0 {
		u := heap.Pop(&pq).(*PriorityQueueItem)
		if u.x == targetX && u.y == targetY {
			break
		}

		for _, d := range directions {
			neighborX := u.x + d[0]
			neighborY := u.y + d[1]
			if neighborX >= 0 && neighborX < g.n && neighborY >= 0 && neighborY < g.n {
				alt := dist[u.x][u.y] + g.matrix[u.x*g.n+u.y][neighborX*g.n+neighborY]
				if alt < dist[neighborX][neighborY] {
					dist[neighborX][neighborY] = alt
					prev[neighborX][neighborY] = u.x*g.n + u.y
					heap.Push(&pq, &PriorityQueueItem{
						x:        neighborX,
						y:        neighborY,
						priority: alt,
					})
				}
			}
		}
	}

	if dist[targetX][targetY] == math.MaxInt32 {
		fmt.Println("No path found")
		return nil, -1
	}

	// Reconstruct the shortest path
	path := [][]int{}
	for u := targetX*g.n + targetY; u != startX*g.n+startY; u = prev[u/g.n][u%g.n] {
		path = append([][]int{{u / g.n, u % g.n}}, path...)
	}
	path = append([][]int{{startX, startY}}, path...)

	return path, dist[targetX][targetY]
}
