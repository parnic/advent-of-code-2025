package days

import (
	"fmt"
	"slices"

	u "parnic.com/aoc2025/utilities"
)

type Day04 struct {
	grid [][]bool
	min  u.Vec2i
	max  u.Vec2i
}

func (d *Day04) Parse() {
	m := u.GetStringLines("04p")
	dim := len(m)
	d.grid = make([][]bool, dim)
	for lineIdx, line := range m {
		d.grid[lineIdx] = make([]bool, dim)
		for col, ch := range line {
			if ch == '@' {
				d.grid[lineIdx][col] = true
			}
		}
	}
	d.min = u.ZeroVec2
	d.max = u.Vec2i{X: dim - 1, Y: dim - 1}
}

func (d Day04) Num() int {
	return 4
}

func (d *Day04) Part1() string {
	var accessible int
	for row := range d.grid {
		for col, val := range d.grid[row] {
			if val {
				pt := u.Vec2i{X: col, Y: row}
				neighbors := pt.GetBoundedNeighbors(d.min, d.max)
				if u.CountFunc(neighbors, func(other u.Vec2i) bool { return d.grid[other.Y][other.X] }) < 4 {
					accessible++
				}
			}
		}
	}
	return fmt.Sprintf("Accessible rolls: %s%d%s", u.TextBold, accessible, u.TextReset)
}

func (d *Day04) cloneGrid(grid [][]bool) [][]bool {
	newGrid := make([][]bool, len(grid))
	for idx, row := range grid {
		newGrid[idx] = slices.Clone(row)
	}

	return newGrid
}

func (d *Day04) Part2() string {
	removed := 0
	currGrid := d.cloneGrid(d.grid)

	for {
		lastGrid := d.cloneGrid(currGrid)
		lastRemoved := removed

		for row := range lastGrid {
			for col, val := range lastGrid[row] {
				if val {
					pt := u.Vec2i{X: col, Y: row}
					neighbors := pt.GetBoundedNeighbors(d.min, d.max)
					if u.CountFunc(neighbors, func(other u.Vec2i) bool { return currGrid[other.Y][other.X] }) < 4 {
						currGrid[row][col] = false
						removed++
					}
				}
			}
		}

		if lastRemoved == removed {
			break
		}
	}

	return fmt.Sprintf("Removable rolls: %s%d%s", u.TextBold, removed, u.TextReset)
}
