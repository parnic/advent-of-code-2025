package days

import (
	"fmt"
	"strconv"
	"strings"

	u "parnic.com/aoc2025/utilities"
)

type Day12_Region struct {
	width         int
	height        int
	presentCounts []int
}

type Day12 struct {
	presentAreas []int
	regions      []Day12_Region
}

func (d *Day12) Parse() {
	lines := u.GetStringLines("12p")
	parsingPresent := false
	for _, line := range lines {
		if !parsingPresent {
			parts := strings.Split(line, ": ")
			if len(parts) == 1 {
				_, err := strconv.Atoi(parts[0][:len(parts[0])-1])
				if err != nil {
					panic(err)
				}
				d.presentAreas = append(d.presentAreas, 0)
				parsingPresent = true
				continue
			}

			region := Day12_Region{}

			areaParts := strings.Split(parts[0], "x")
			width, err := strconv.Atoi(areaParts[0])
			if err != nil {
				panic(err)
			}
			region.width = width

			height, err := strconv.Atoi(areaParts[1])
			if err != nil {
				panic(err)
			}
			region.height = height

			presentCounts := strings.Fields(parts[1])
			region.presentCounts = make([]int, 0, len(presentCounts))
			for _, count := range presentCounts {
				presentCount, err := strconv.Atoi(count)
				if err != nil {
					panic(err)
				}
				region.presentCounts = append(region.presentCounts, presentCount)
			}

			d.regions = append(d.regions, region)
			continue
		}
		if len(line) == 0 {
			parsingPresent = false
			continue
		}

		lineArea := strings.Count(line, "#")
		d.presentAreas[len(d.presentAreas)-1] += lineArea
	}
}

func (d Day12) Num() int {
	return 12
}

func (d *Day12) Part1() string {
	var fitting int
	for _, region := range d.regions {
		regionArea := region.width * region.height
		var requiredArea int
		for idx, presentCount := range region.presentCounts {
			requiredArea += d.presentAreas[idx] * presentCount
		}
		if requiredArea <= regionArea {
			fitting++
		}
	}
	return fmt.Sprintf("# regions that can fit all presents: %s%d%s", u.TextBold, fitting, u.TextReset)
}

func (d *Day12) Part2() string {
	return fmt.Sprintf("%sM%se%sr%sr%sy %sC%sh%sr%si%ss%st%sm%sa%ss%s!%s",
		u.ColorBrightRed,
		u.ColorBrightGreen,
		u.ColorBrightRed,
		u.ColorBrightGreen,
		u.ColorBrightRed,
		u.ColorBrightGreen,
		u.ColorBrightRed,
		u.ColorBrightGreen,
		u.ColorBrightRed,
		u.ColorBrightGreen,
		u.ColorBrightRed,
		u.ColorBrightGreen,
		u.ColorBrightRed,
		u.ColorBrightGreen,
		u.ColorBrightRed,
		u.ColorBrightGreen,
	)
}
