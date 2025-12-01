package days

import (
	"fmt"
	"strconv"

	"parnic.com/aoc2025/utilities"
)

type Day01 struct {
	instructions []string
}

func (d *Day01) Parse() {
	d.instructions = utilities.GetStringLines("01p")
}

func (d Day01) Num() int {
	return 1
}

func (d *Day01) Part1() string {
	var password int64
	curr := int64(50)
	for _, inst := range d.instructions {
		dir := inst[0]
		amt, err := strconv.ParseInt(inst[1:], 10, 64)
		if err != nil {
			panic(err)
		}
		if dir == 'L' {
			amt = -amt
		}

		curr += amt
		for curr < 0 {
			curr += 100
		}
		curr = curr % 100
		if curr == 0 {
			password++
		}
	}

	return fmt.Sprintf("Password: %s%d%s", utilities.TextBold, password, utilities.TextReset)
}

func (d *Day01) Part2() string {
	var password int64
	curr := int64(50)
	for _, inst := range d.instructions {
		dir := inst[0]
		amt, err := strconv.ParseInt(inst[1:], 10, 64)
		if err != nil {
			panic(err)
		}
		if dir == 'L' {
			amt = -amt
		}

		// dumb version because i can't get the smart version working quite right and what's 1ms between friends?
		for range amt {
			if dir == 'L' {
				curr--
			} else {
				curr++
			}
			if curr == 100 {
				curr = 0
			}
			switch curr {
			case 0:
				password++
			case -1:
				curr = 99
			}
		}
	}

	return fmt.Sprintf("Password: %s%d%s", utilities.TextBold, password, utilities.TextReset)
}
