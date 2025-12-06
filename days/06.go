package days

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	u "parnic.com/aoc2025/utilities"
)

type Day06_Problem struct {
	nums []int64
	op   rune
}

type Day06 struct {
	problems     []Day06_Problem
	cephProblems []Day06_Problem
}

func (d *Day06) Parse() {
	lines := u.GetStringLines("06p")
	r := regexp.MustCompile(`\s+`)
	line1 := strings.TrimSpace(r.ReplaceAllString(lines[0], " "))
	d.problems = make([]Day06_Problem, strings.Count(line1, " ")+1)
	numVals := len(lines) - 1
	opIdx := len(lines) - 1
	// parse as normal human math
	for i, line := range lines {
		for pi, str := range strings.Split(strings.TrimSpace(r.ReplaceAllString(line, " ")), " ") {
			if i == opIdx {
				d.problems[pi].op = rune(str[0])
			} else {
				if i == 0 {
					d.problems[pi].nums = make([]int64, numVals)
				}
				num, _ := strconv.ParseInt(str, 10, 32)
				d.problems[pi].nums[i] = num
			}
		}
	}

	d.cephProblems = make([]Day06_Problem, strings.Count(line1, " ")+1)
	// parse as jacked-up cephalopod math
	lineLen := len(lines[0])
	problemIdx := len(d.problems) - 1
	for i := lineLen - 1; i >= 0; i-- {
		divider := true
		sb := strings.Builder{}
		for j := 0; j < len(lines)-1; j++ {
			inspect := lines[j][i]
			if inspect != ' ' {
				divider = false
			}
			if inspect == ' ' || inspect == '+' || inspect == '*' {
				continue
			}

			sb.WriteByte(inspect)
		}

		if divider {
			problemIdx--
		} else {
			num, err := strconv.ParseInt(sb.String(), 10, 64)
			if err != nil {
				panic(err)
			}
			d.cephProblems[problemIdx].nums = append(d.cephProblems[problemIdx].nums, num)
			d.cephProblems[problemIdx].op = d.problems[problemIdx].op
		}
	}
}

func (d Day06) Num() int {
	return 6
}

func (d *Day06) solve(problems []Day06_Problem) int64 {
	var sum int64
	for _, problem := range problems {
		var total int64
		for i, num := range problem.nums {
			if i == 0 {
				total = int64(num)
				continue
			}

			switch problem.op {
			case '*':
				total *= num
			case '+':
				total += num
			}
		}

		sum += total
	}
	return sum
}

func (d *Day06) Part1() string {
	sum := d.solve(d.problems)
	return fmt.Sprintf("Human total: %s%d%s", u.TextBold, sum, u.TextReset)
}

func (d *Day06) Part2() string {
	sum := d.solve(d.cephProblems)
	return fmt.Sprintf("Cephalopod total: %s%d%s", u.TextBold, sum, u.TextReset)
}
