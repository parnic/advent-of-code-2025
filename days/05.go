package days

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"

	u "parnic.com/aoc2025/utilities"
)

type Day05_IDRange struct {
	min uint64
	max uint64
}

type Day05 struct {
	freshRanges []Day05_IDRange
	ids         []uint64
}

func (d *Day05) Parse() {
	d.freshRanges = make([]Day05_IDRange, 0)
	lines := u.GetStringLines("05p")
	section := 0
	for _, line := range lines {
		if len(line) == 0 {
			section++
			continue
		}

		switch section {
		case 0:
			parts := strings.Split(line, "-")
			r := Day05_IDRange{}
			r.min, _ = strconv.ParseUint(parts[0], 10, 64)
			r.max, _ = strconv.ParseUint(parts[1], 10, 64)
			d.freshRanges = append(d.freshRanges, r)

		default:
			id, _ := strconv.ParseUint(line, 10, 64)
			d.ids = append(d.ids, id)
		}
	}
}

func (d Day05) Num() int {
	return 5
}

func (d *Day05) Part1() string {
	var numFresh uint64
	for _, id := range d.ids {
		for _, r := range d.freshRanges {
			if r.min <= id && r.max >= id {
				numFresh++
				break
			}
		}
	}
	return fmt.Sprintf("Num fresh ingredients: %s%d%s", u.TextBold, numFresh, u.TextReset)
}

func (d *Day05) Part2() string {
	ranges := slices.Clone(d.freshRanges)
	slices.SortFunc(ranges, func(r1 Day05_IDRange, r2 Day05_IDRange) int { return cmp.Compare(r1.min, r2.min) })
	result := make([]Day05_IDRange, 0, len(ranges))
	// scan ahead and subsume any ranges that overlap ours
	for i := 0; i < len(ranges); {
		j := i + 1
		for len(ranges) > j && ranges[j].min <= ranges[i].max {
			if ranges[j].max > ranges[i].max {
				ranges[i].max = ranges[j].max
			}
			j++
		}
		result = append(result, ranges[i])
		i = j
	}

	var numFresh uint64
	for _, r := range result {
		numFresh += r.max - r.min + 1
	}
	return fmt.Sprintf("Total fresh ingredient count: %s%d%s", u.TextBold, numFresh, u.TextReset)
}
