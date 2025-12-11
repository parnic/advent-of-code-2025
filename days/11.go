package days

import (
	"fmt"
	"strings"

	u "parnic.com/aoc2025/utilities"
)

type Day11 struct {
	paths map[string][]string
}

func (d *Day11) Parse() {
	d.paths = make(map[string][]string)
	lines := u.GetStringLines("11p")
	for _, line := range lines {
		dev := strings.Split(line, ": ")
		cons := strings.Fields(dev[1])
		d.paths[dev[0]] = cons
	}
}

func (d Day11) Num() int {
	return 11
}

func (d *Day11) numExits(start string, stop string, memo map[string]int64) int64 {
	if start == stop {
		return 1
	}
	if _, exists := memo[start]; exists {
		return memo[start]
	}

	var exits int64
	for _, candidate := range d.paths[start] {
		if candidateExits := d.numExits(candidate, stop, memo); candidateExits > 0 {
			exits += candidateExits
		}
	}

	memo[start] = exits
	return exits
}

func (d *Day11) Part1() string {
	memo := make(map[string]int64)
	d.numExits("you", "out", memo)
	return fmt.Sprintf("Paths from you to out: %s%d%s", u.TextBold, memo["you"], u.TextReset)
}

func (d *Day11) Part2() string {
	// map out the individual routes and determine which ones actually have a throughline to the end
	memo := make(map[string]int64)
	svrToDac := d.numExits("svr", "dac", memo)
	clear(memo)
	svrToFft := d.numExits("svr", "fft", memo)
	clear(memo)
	dacToFft := d.numExits("dac", "fft", memo)
	clear(memo)
	fftToDac := d.numExits("fft", "dac", memo)
	clear(memo)
	fftToOut := d.numExits("fft", "out", memo)
	clear(memo)
	dacToOut := d.numExits("dac", "out", memo)

	svrThruDac := svrToDac * dacToFft * fftToOut
	svrThruFft := svrToFft * fftToDac * dacToOut
	svrThruDacAndFft := svrThruDac + svrThruFft

	return fmt.Sprintf("Paths from svr to out through dac + fft: %s%d%s", u.TextBold, svrThruDacAndFft, u.TextReset)
}
