package days

import (
	"fmt"
	"maps"

	u "parnic.com/aoc2025/utilities"
)

type Day07 struct {
	start     u.Vec2i
	width     int
	height    int
	splitters []u.Vec2i
}

func (d *Day07) Parse() {
	lines := u.GetStringLines("07p")
	d.width = len(lines[0])
	d.height = len(lines)
	for row, rowLine := range lines {
		for col, colChar := range rowLine {
			switch colChar {
			case 'S':
				d.start.X = col
				d.start.Y = row

			case '^':
				splitter := u.Vec2i{X: col, Y: row}
				d.splitters = append(d.splitters, splitter)
			}
		}
	}
}

func (d Day07) Num() int {
	return 7
}

func (d *Day07) Part1() string {
	var numSplits int
	beamLocs := make(map[u.Vec2i]struct{})
	beamLocs[d.start] = struct{}{}
	movingBeams := maps.Clone(beamLocs)

	processing := true
	for processing {
		processing = false

		lastBeams := maps.Clone(movingBeams)
		for beam := range lastBeams {
			delete(movingBeams, beam)
			beam = beam.AddVec(u.Down)
			if beam.Y > d.height || beam.X < 0 || beam.X > d.width {
				continue
			}

			processing = true

			if u.ArrayContains(d.splitters, beam) {
				numSplits++
				left := beam.AddVec(u.Left)
				right := beam.AddVec(u.Right)
				if _, ok := beamLocs[left]; !ok {
					beamLocs[left] = struct{}{}
					movingBeams[left] = struct{}{}
				}
				if _, ok := beamLocs[right]; !ok {
					beamLocs[right] = struct{}{}
					movingBeams[right] = struct{}{}
				}
				continue
			}

			movingBeams[beam] = struct{}{}
		}
	}
	return fmt.Sprintf("Beam will split %s%d%s times", u.TextBold, numSplits, u.TextReset)
}

func (d *Day07) memoizeGetTimelines(memo map[u.Vec2i]uint64, p u.Vec2i) uint64 {
	if p.Y == d.height {
		return 1
	}
	if cached, ok := memo[p]; ok {
		return cached
	}

	var timelines uint64
	if u.ArrayContains(d.splitters, p) {
		if p.X > 0 {
			timelines = d.memoizeGetTimelines(memo, p.AddVec(u.Left))
		}
		if p.X < d.width {
			timelines += d.memoizeGetTimelines(memo, p.AddVec(u.Right))
		}
	} else {
		timelines = d.memoizeGetTimelines(memo, p.AddVec(u.Down))
	}

	memo[p] = timelines
	return timelines
}

func (d *Day07) Part2() string {
	memo := make(map[u.Vec2i]uint64)
	timelines := d.memoizeGetTimelines(memo, d.start)
	return fmt.Sprintf("Timelines for 1 particle: %s%d%s", u.TextBold, timelines, u.TextReset)
}
