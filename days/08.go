package days

import (
	"cmp"
	"fmt"
	"slices"

	u "parnic.com/aoc2025/utilities"
)

type Day08 struct {
	boxes []u.Vec3[int]
}

type Day08_Distance struct {
	a    u.Vec3[int]
	b    u.Vec3[int]
	dist int
}

func (d *Day08) Parse() {
	lines := u.GetStringLines("08p")
	d.boxes = make([]u.Vec3[int], len(lines))
	for i, line := range lines {
		var err error
		d.boxes[i], err = u.ParseVec3[int](line)
		if err != nil {
			panic(err)
		}
	}
}

func (d Day08) Num() int {
	return 8
}

func (d *Day08) ConnectBoxes(circuits [][]u.Vec3[int], a u.Vec3[int], b u.Vec3[int]) [][]u.Vec3[int] {
	ai := slices.IndexFunc(circuits, func(list []u.Vec3[int]) bool {
		return u.ArrayContains(list, a)
	})
	bi := slices.IndexFunc(circuits, func(list []u.Vec3[int]) bool {
		return u.ArrayContains(list, b)
	})
	// neither is connected to anything
	if ai < 0 && bi < 0 {
		circuits = append(circuits, []u.Vec3[int]{a, b})
		return circuits
	}
	// they're both in the same circuit already
	if ai == bi {
		return circuits
	}

	// a is in a circuit, b is not
	if ai >= 0 && bi < 0 {
		circuits[ai] = append(circuits[ai], b)
		return circuits
	}

	// b is in a circuit, a is not
	if ai < 0 && bi >= 0 {
		circuits[bi] = append(circuits[bi], a)
		return circuits
	}

	// they're both in separate circuits, so we need to merge them.
	from := bi
	to := ai
	if from < to {
		from, to = to, from
	}
	bOthers := circuits[from]
	circuits = slices.Delete(circuits, from, from+1)
	circuits[to] = append(circuits[to], bOthers...)
	return circuits
}

func (d *Day08) buildCircuits(stopCondition func(circuits [][]u.Vec3[int], idx int) bool) ([][]u.Vec3[int], u.Vec3[int], u.Vec3[int]) {
	distances := make([]Day08_Distance, 0)
	for i := range d.boxes {
		for j := i + 1; j < len(d.boxes); j++ {
			distances = append(distances, Day08_Distance{
				a:    d.boxes[i],
				b:    d.boxes[j],
				dist: d.boxes[i].DistanceSquared(d.boxes[j]),
			})
		}
	}
	slices.SortFunc(distances, func(a Day08_Distance, b Day08_Distance) int {
		return cmp.Compare(a.dist, b.dist)
	})

	circuits := make([][]u.Vec3[int], 0)
	for _, b := range d.boxes {
		singleList := []u.Vec3[int]{b}
		circuits = append(circuits, singleList)
	}
	i := 0
	for i = range distances {
		circuits = d.ConnectBoxes(circuits, distances[i].a, distances[i].b)
		if stopCondition(circuits, i) {
			break
		}
	}

	return circuits, distances[i].a, distances[i].b
}

func (d *Day08) Part1() string {
	circuits, _, _ := d.buildCircuits(func(circuits [][]u.Vec3[int], idx int) bool {
		return idx == 999
	})

	slices.SortFunc(circuits, func(a []u.Vec3[int], b []u.Vec3[int]) int {
		return cmp.Compare(len(b), len(a))
	})
	sizeMultiple := len(circuits[0]) * len(circuits[1]) * len(circuits[2])

	return fmt.Sprintf("Three largest circuit value: %s%d%s", u.TextBold, sizeMultiple, u.TextReset)
}

func (d *Day08) Part2() string {
	_, lastA, lastB := d.buildCircuits(func(circuits [][]u.Vec3[int], idx int) bool {
		return len(circuits) == 1
	})

	ans := lastA.X * lastB.X
	return fmt.Sprintf("Final connected circuit value: %s%d%s", u.TextBold, ans, u.TextReset)
}
