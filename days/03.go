package days

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	u "parnic.com/aoc2025/utilities"
)

type Day03 struct {
	lines []string
}

func (d *Day03) Parse() {
	d.lines = u.GetStringLines("03p")
}

func (d Day03) Num() int {
	return 3
}

func (d *Day03) Part1() string {
	var sum int64
	// my first thought. sort the list, pick out the largest two numbers we find in order, combine them.
	for _, line := range d.lines {
		iLine := make([]int, len(line))
		for i := range line {
			iLine[i] = int(line[i] - '0')
		}
		firstSorted := slices.Sorted(slices.Values(iLine[:len(iLine)-1]))
		firstLargest := firstSorted[len(firstSorted)-1]
		idx := slices.Index(iLine, firstLargest)
		secondSorted := slices.Sorted(slices.Values(iLine[idx+1:]))
		secondLargest := secondSorted[len(secondSorted)-1]
		numStr := fmt.Sprintf("%d%d", firstLargest, secondLargest)
		num, _ := strconv.ParseInt(numStr, 10, 64)
		sum += num
	}
	return fmt.Sprintf("Total output joltage from 2 batteries: %s%d%s", u.TextBold, sum, u.TextReset)
}

func (d *Day03) Part2() string {
	var sum int64
	numBatteries := 12

	// first approach clearly won't work for 12 batteries.
	// this is similar, but just extracts the 12 largest values we find in order, which seems to solve the problem.
	// i kinda feel like it shouldn't necessarily be that easy, but...here we are.
	for _, text := range d.lines {
		indexes := make([]int, numBatteries)
		f := func(idx int) int {
			largestVal := byte('0')
			largestIdx := len(text) - numBatteries - idx
			startIdx := 0
			if idx > 0 {
				startIdx = indexes[idx-1] + 1
			}
			for i := startIdx; i < len(text)-(numBatteries-1-idx); i++ {
				if text[i] > largestVal {
					largestVal = text[i]
					largestIdx = i
				}
			}

			indexes[idx] = largestIdx
			return largestIdx
		}

		sb := strings.Builder{}
		for idx := range indexes {
			i := f(idx)
			sb.WriteByte(text[i])
		}

		n, _ := strconv.ParseInt(sb.String(), 10, 64)
		sum += n
	}

	return fmt.Sprintf("Total output joltage from 12 batteries: %s%d%s", u.TextBold, sum, u.TextReset)
}
