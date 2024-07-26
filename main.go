package main

import (
	"fmt"
)

func main() {
	size := 5
	player := Player{}
	myMap := newMap("MyMap", player, size)
	snake := newSnake()
	snake.x = 0
	snake.y = 0

	targetX, targetY := 4, 4

	// Print the map details
	fmt.Printf("Map Name: %s\n", myMap.name)
	fmt.Printf("Map Size: %d\n", myMap.size)
	fmt.Printf("Player Name: %s\n", myMap.player.name)
	fmt.Printf("Player Points: %d\n", myMap.player.points)
	fmt.Printf("Player Position: (%d, %d)\n", myMap.player.x, myMap.player.y)
	fmt.Printf("Snake Position: (%d, %d)\n", myMap.snake.x, myMap.snake.y)

	// Print the graph
	fmt.Println("Graph Adjacency Matrix:")
	for i := 0; i < myMap.size; i++ {
		for j := 0; j < myMap.size; j++ {
			fmt.Printf("%d ", myMap.web.matrix[i][j])
		}
		fmt.Println()
	}

	// Ensure target coordinates are within bounds
	if targetX < 0 || targetX >= size || targetY < 0 || targetY >= size {
		fmt.Println("Target position is out of bounds")
		return
	}

	// Find the shortest path
	path, distance := myMap.web.shortestPath(snake.x, snake.y, targetX, targetY)
	if distance == -1 {
		fmt.Println("No path found")
	} else {
		fmt.Printf("Shortest path: %v\n", path)
		fmt.Printf("Total distance: %d\n", distance)
	}
}
