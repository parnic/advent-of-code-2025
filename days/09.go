package days

import (
	"fmt"

	u "parnic.com/aoc2025/utilities"
)

type Day09 struct {
	points []u.Vec2[int]
	rects  []u.Rectangle[int]
	edges  []u.Rectangle[int] // treat edges as rectangles for easy overlap checks
}

func (d *Day09) Parse() {
	lines := u.GetStringLines("09p")
	numLines := len(lines)

	d.points = make([]u.Vec2[int], 0, numLines)
	d.rects = make([]u.Rectangle[int], 0, numLines*(numLines/2))
	d.edges = make([]u.Rectangle[int], 0, numLines)

	addRect := func(a u.Vec2[int], b u.Vec2[int], list []u.Rectangle[int]) []u.Rectangle[int] {
		rect := u.Rectangle[int]{Min: a, Max: b}.Canonical()
		rect.Max = rect.Max.AddVec(u.IVec2One)
		list = append(list, rect)
		return list
	}

	for i, line := range lines {
		p, err := u.ParseVec2[int](line)
		if err != nil {
			panic(err)
		}

		for _, prev := range d.points {
			d.rects = addRect(p, prev, d.rects)
		}
		d.points = append(d.points, p)
		if len(d.points) > 1 {
			d.edges = addRect(d.points[i-1], d.points[i], d.edges)
		}
	}
	// connect the end to the front
	d.edges = addRect(d.points[len(d.points)-1], d.points[0], d.edges)
}

func (d Day09) Num() int {
	return 9
}

func (d *Day09) Part1() string {
	var largestArea int
	for _, r := range d.rects {
		area := r.Width() * r.Height()
		largestArea = max(largestArea, area)
	}
	return fmt.Sprintf("Largest possible rectangle: %s%d%s", u.TextBold, largestArea, u.TextReset)
}

func (d *Day09) Part2() string {
	var largestArea int

	for _, r := range d.rects {
		invalid := false
		area := r.Width() * r.Height()

		for _, e := range d.edges {
			// if a contracted rectangle overlaps any edges, it extends outside the bounds
			if e.Overlaps(r.Inset(1)) {
				invalid = true
				break
			}
		}

		if !invalid {
			largestArea = max(largestArea, area)
		}
	}

	return fmt.Sprintf("Largest constrained rectangle: %s%d%s", u.TextBold, largestArea, u.TextReset)
}
