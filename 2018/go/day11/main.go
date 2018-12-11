package main

import "log"

func main() {
	grid := computerPowerLevels(5153)
	x, y, _ := calculateSquare(grid, 3)

	log.Println(x, ",", y)

	var max, a, b, s int
	for i := 1; i <= 300; i++ {
		log.Println(i)
		x, y, power := calculateSquare(grid, i)
		if power > max {
			max = power
			a = x
			b = y
			s = i
		}
	}

	log.Println(a, b, s)
}

func calculateSquare(grid [][]int, size int) (int, int, int) {
	var x, y, sum int
	imax := len(grid) - size
	for i := range grid {
		if i >= imax {
			continue
		}
		jmax := len(grid[i]) - size
		for j := range grid[i] {
			if j >= jmax {
				continue
			}
			var thisSum int
			for a := i; a < i+size; a++ {
				for b := j; b < j+size; b++ {
					thisSum += grid[a][b]
				}
			}
			// thisSum := grid[i][j] + grid[i+1][j] + grid[i+2][j] + grid[i][j+1] + grid[i][j+2] + grid[i+1][j+1] + grid[i+1][j+2] + grid[i+2][j] + grid[i+2][j+1] + grid[i+2][j+2]
			if thisSum > sum {
				x, y = i+1, j+1
				sum = thisSum
			}
		}
	}

	return x, y, sum
}

func computerPowerLevels(serial int) [][]int {
	grid := make([][]int, 300)

	for i := range grid {
		grid[i] = make([]int, 300)
		for j := range grid[i] {
			grid[i][j] = calculatePower(i+1, j+1, serial)
		}
	}

	return grid
}

func calculatePower(x, y, serial int) int {
	rackID := x + 10
	return (((((rackID * y) + serial) * rackID) % 1000) / 100) - 5
}
