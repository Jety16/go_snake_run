# Go Snake Run

A simple Go project to create a snake game with a player and a snake on a map. The project includes functionalities to find the shortest path using Dijkstra's algorithm, move the snake, and allow the player to move.

## Features

- [x] **Make the Kickstarter for my project**
- [x]  **Add method to find the shortest path from point A to point B and return the distance**
- [x] **Use the method with the snake and player**
- [ ]  **Update the map weights based on that update**
- [ ] **Make the snake move to the first node from the shortest path**
- [ ] **Make the player able to move**


### Project Structure

- `main.go`: Entry point of the application.
- `player.go`: Contains the `Player` struct and its methods.
- `map.go`: Contains the `Graph` and `Map` structs and methods, including Dijkstra's algorithm.
- `snake.go`: Contains the `Snake` struct and its methods.
- `priority_queue.go`: Implements a simple priority queue for Dijkstra's algorithm.
