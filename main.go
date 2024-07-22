package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Create a player
	player := newPlayer("John", 3, 2)

	// Create a new map
	gameMap := newMap("MyGameMap", *player, 5)

	// Print the map details
	fmt.Printf("Map Name: %s\n", gameMap.name)
	fmt.Printf("Map Size: %d\n", gameMap.size)
	fmt.Printf("Player Name: %s\n", gameMap.player.name)
	fmt.Printf("Player Points: %d\n", gameMap.player.points)
	fmt.Printf("Player Position: (%d, %d)\n", gameMap.player.x, gameMap.player.y)
	fmt.Printf("Snake Position: (%d, %d)\n", gameMap.snake.x, gameMap.snake.y)

	// Print the graph
	fmt.Println("Graph Adjacency Matrix:")
	for i := 0; i < gameMap.size; i++ {
		for j := 0; j < gameMap.size; j++ {
			fmt.Printf("%d ", gameMap.web.matrix[i][j])
		}
		fmt.Println()
	}
}
