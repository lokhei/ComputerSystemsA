package main

const alive = 255
const dead = 0

func mod(x, m int) int {
	return (x + m) % m
}

//calculates number of neighbours of cell
func calculateNeighbours(p golParams, x, y int, world [][]byte) int {
	neighbours := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i != 0 || j != 0 { //not [y][x]
				if world[mod(y+i, p.imageHeight)][mod(x+j, p.imageWidth)] == alive {
					neighbours++
				}
			}
		}
	}
	return neighbours
}

//takes the current state of the world and completes one evolution of the world. It then returns the result.
func calculateNextState(p golParams, world [][]byte) [][]byte {
	//makes a new world
	newWorld := make([][]byte, p.imageHeight)
	for i := range newWorld {
		newWorld[i] = make([]byte, p.imageWidth)
	}
	//sets cells to dead or alive according to num of neighbours
	for y := 0; y < p.imageHeight; y++ {
		for x := 0; x < p.imageWidth; x++ {
			neighbours := calculateNeighbours(p, x, y, world)
			if world[y][x] == alive {
				if neighbours == 2 || neighbours == 3 {
					newWorld[y][x] = alive
				} else {
					newWorld[y][x] = dead
				}
			} else {
				if neighbours == 3 {
					newWorld[y][x] = alive
				} else {
					newWorld[y][x] = dead
				}
			}
		}
	}
	return newWorld
}

//takes the world as input and returns the (x, y) coordinates of all the cells that are alive.
func calculateAliveCells(p golParams, world [][]byte) []cell {
	aliveCells := []cell{}

	for y := 0; y < p.imageHeight; y++ {
		for x := 0; x < p.imageWidth; x++ {
			if world[y][x] == alive {
				aliveCells = append(aliveCells, cell{x: x, y: y})
			}
		}
	}

	return aliveCells
}
