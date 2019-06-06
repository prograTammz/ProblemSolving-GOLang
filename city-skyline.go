package main

import (
	"math"
)

func main() {

}
func maxIncreaseKeepingSkyline(grid [][]int) int {
	maxRowCol := Maxcolrow(grid)
	length := len(grid[0])
	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < length; j++ {
			result = result + (int(math.Min(float64(maxRowCol[0][i]), float64(maxRowCol[i][j]))) - grid[i][j])
		}
	}
	return result
}
func maxOfRows(grid [][]int) []int {
	var maxRow []int
	var length = len(grid[0])

	for i := 0; i < len(grid); i++ {
		max := 0
		for j := 0; j < length; j++ {
			val := grid[i][j]
			if val > max {
				max = val
			}
		}
		maxRow = append(maxRow, max)
	}
	return maxRow
}
func maxOfCols(grid [][]int) []int {
	var maxCol []int
	var length = len(grid[0])
	for i := 0; i < len(grid); i++ {
		max := 0
		for j := 0; j < length; j++ {
			val := grid[j][i]
			if val > max {
				max = val
			}
		}
		maxCol = append(maxCol, max)
	}
	return maxCol
}
func Maxcolrow(grid [][]int) [][]int {
	maxRowCol := make([][]int, 5)
	var length = len(grid[0])
	for i := 0; i < len(grid); i++ {
		maxRow := 0
		maxCol := 0
		for j := 0; j < length; j++ {
			valRow := grid[i][j]
			valCol := grid[j][i]
			if valRow > maxRow {
				maxRow = valRow
			}
			if valCol > maxCol {
				maxCol = valCol
			}
		}
		maxRowCol[0] = append(maxRowCol[0], maxRow)
		maxRowCol[1] = append(maxRowCol[1], maxCol)
	}
	return maxRowCol
}
