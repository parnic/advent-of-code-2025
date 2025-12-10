package days

import (
	"fmt"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"

	// "github.com/aclements/go-z3/z3"
	u "parnic.com/aoc2025/utilities"
)

type Day10_Machine struct {
	lights  []bool
	buttons [][]int
	joltage []int
}

type Day10 struct {
	machines []Day10_Machine
}

func (d *Day10) Parse() {
	lines := u.GetStringLines("10p")
	d.machines = make([]Day10_Machine, len(lines))
	for i, line := range lines {
		machine := Day10_Machine{}
		fields := strings.Fields(line)
		machine.buttons = make([][]int, len(fields)-2)
		for f, field := range fields {
			switch field[0] {
			case '[':
				machine.lights = make([]bool, len(field)-2)
				for i := 1; i < len(field)-1; i++ {
					if field[i] == '#' {
						machine.lights[i-1] = true
					}
				}

			case '(':
				buttons := strings.Split(field[1:len(field)-1], ",")
				machine.buttons[f-1] = u.Map(buttons, func(btn string) int {
					n, err := strconv.ParseInt(btn, 10, 32)
					if err != nil {
						panic(err)
					}
					return int(n)
				})

			case '{':
				joltages := strings.Split(field[1:len(field)-1], ",")
				machine.joltage = u.Map(joltages, func(joltage string) int {
					n, err := strconv.ParseInt(joltage, 10, 32)
					if err != nil {
						panic(err)
					}
					return int(n)
				})
			}
		}
		d.machines[i] = machine
	}
}

func (d Day10) Num() int {
	return 10
}

func (d *Day10) Part1() string {
	var totalPresses int64
	for m, machine := range d.machines {
		// grab all possible combinations of pressing buttons at most one time
		combinations := make([][]int, 0)
		for i := range machine.buttons {
			u.Combinations(len(machine.buttons), i+1, func(comb []int) {
				combinations = append(combinations, slices.Clone(comb))
			})
		}

		// combinations are in increasing order, so find the first time pressing the given list of buttons results in the target state
		target := machine.lights
		requiredPresses := math.MaxInt
		for _, combination := range combinations {
			state := make([]bool, len(target))
			for _, combo := range combination {
				for _, btn := range machine.buttons[combo] {
					state[btn] = !state[btn]
				}
			}
			if slices.Equal(state, target) {
				requiredPresses = min(requiredPresses, len(combination))
				break
			}
		}

		if requiredPresses == math.MaxInt {
			log.Fatalln("no solution found for machine", m+1)
		}

		totalPresses += int64(requiredPresses)
	}
	return fmt.Sprintf("Total minimum number of presses required: %s%d%s", u.TextBold, totalPresses, u.TextReset)
}

func (d *Day10) Part2() string {
	// var totalPresses int64
	// for _, machine := range d.machines {
	// 	presses := d.calculateMinButtonJoltagePresses(machine)
	// 	totalPresses += int64(presses)
	// }
	// return fmt.Sprintf("Total minimum number of presses for joltage: %s%d%s", u.TextBold, totalPresses, u.TextReset)
	return fmt.Sprintf("Uncomment this file and install z3 to get your answer.")
}

// func (d *Day10) calculateMinButtonJoltagePresses(machine Day10_Machine) int {
// 	numButtons := len(machine.buttons)

// 	ctx := z3.NewContext(nil)
// 	solver := z3.NewSolver(ctx)

// 	intSort := ctx.IntSort()
// 	zero := ctx.FromInt(0, intSort).(z3.Int)
// 	one := ctx.FromInt(1, intSort).(z3.Int)

// 	buttons := make([]z3.Int, numButtons)
// 	for i := 0; i < numButtons; i++ {
// 		buttons[i] = ctx.IntConst("button_" + strconv.Itoa(i))
// 		solver.Assert(buttons[i].GE(zero))
// 	}

// 	for counterIdx, targetValue := range machine.joltage {
// 		var buttonsThatIncrement []z3.Int
// 		for buttonIdx, button := range machine.buttons {
// 			for _, affectedCounter := range button {
// 				if affectedCounter == counterIdx {
// 					buttonsThatIncrement = append(buttonsThatIncrement, buttons[buttonIdx])
// 					break
// 				}
// 			}
// 		}

// 		rhs := ctx.FromInt(int64(targetValue), intSort).(z3.Int)

// 		if len(buttonsThatIncrement) == 0 {
// 			if targetValue > 0 {
// 				solver.Assert(zero.Eq(one))
// 			}
// 		} else {
// 			sum := buttonsThatIncrement[0]
// 			for _, t := range buttonsThatIncrement[1:] {
// 				sum = sum.Add(t)
// 			}
// 			solver.Assert(sum.Eq(rhs))
// 		}
// 	}

// 	tot := ctx.IntConst("total")
// 	if len(buttons) > 0 {
// 		sumAll := buttons[0]
// 		for _, x := range buttons[1:] {
// 			sumAll = sumAll.Add(x)
// 		}
// 		solver.Assert(tot.Eq(sumAll))
// 	} else {
// 		solver.Assert(tot.Eq(zero))
// 	}

// 	minResult := -1
// 	for {
// 		sat, err := solver.Check()
// 		if !sat || err != nil {
// 			break
// 		}

// 		model := solver.Model()
// 		res := model.Eval(tot, true)
// 		val, _, _ := res.(z3.Int).AsInt64()

// 		minResult = int(val)

// 		cur := ctx.FromInt(val, intSort).(z3.Int)
// 		solver.Assert(tot.LT(cur))
// 	}

// 	return minResult
// }
