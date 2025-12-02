package days

import (
	"fmt"
	"strconv"
	"strings"

	u "parnic.com/aoc2025/utilities"
)

type Day02 struct {
	ranges []string
}

func (d *Day02) Parse() {
	data := u.GetStringContents("02p")
	d.ranges = strings.Split(data, ",")
}

func (d Day02) Num() int {
	return 2
}

func (d *Day02) Part1() string {
	var sum uint64
	for _, idrange := range d.ranges {
		vals := strings.Split(idrange, "-")
		min, _ := strconv.ParseUint(vals[0], 10, 64)
		max, _ := strconv.ParseUint(vals[1], 10, 64)

		for id := min; id <= max; id++ {
			idstr := strconv.FormatUint(id, 10)
			for i := 1; i < len(idstr); i++ {
				if idstr[:i] == idstr[i:] {
					sum += id
				}
			}
		}
	}
	return fmt.Sprintf("Invalid ID sum: %s%d%s", u.TextBold, sum, u.TextReset)
}

func (d *Day02) Part2() string {
	var sum uint64
	for _, idrange := range d.ranges {
		vals := strings.Split(idrange, "-")
		min, _ := strconv.ParseUint(vals[0], 10, 64)
		max, _ := strconv.ParseUint(vals[1], 10, 64)

		for id := min; id <= max; id++ {
			idstr := strconv.FormatUint(id, 10)
			for i := 1; i < len(idstr); i++ {
				chunked := u.ChunkString(idstr, i)
				if chunked == nil {
					continue
				}
				matches := u.AllFunc(chunked, func(c string) bool { return c == chunked[0] })
				if matches {
					sum += id
					break
				}
			}
		}
	}
	return fmt.Sprintf("Invalid ID sum: %s%d%s", u.TextBold, sum, u.TextReset)
}
