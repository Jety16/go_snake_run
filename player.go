package main

// Player struct definition
type Player struct {
	name   string
	points int
	x      int
	y      int
}

// Function to create a new Player
func newPlayer(name string, x int, y int) *Player {
	return &Player{
		name:   name,
		points: 0, // Initial points are set to 0
		x:      x,
		y:      y,
	}
}
