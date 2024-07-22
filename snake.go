package main

type Snake struct {
	x int
	y int
}

func newSnake() *Snake {
	return &Snake{
		x: 0,
		y: 0,
	}
}
