package main

// Snake struct definition represents the snake's position on the map.
type Snake struct {
	x int
	y int
}

// newSnake creates a new Snake instance with initial position (0, 0).
func newSnake() *Snake {
	return &Snake{
		x: 0,
		y: 0,
	}
}
