package utilities

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"parnic.com/aoc2025/inputs"
)

func downloadInput(filename string) error {
	if len(filename) < 2 {
		return fmt.Errorf("unexpected filename %s", filename)
	}

	dayStr := filename[:2]
	day, err := strconv.ParseUint(dayStr, 10, 32)
	if err != nil {
		return fmt.Errorf("failed to parse filename %s as day: %w", filename, err)
	}

	input, err := downloadAOCURL(fmt.Sprintf("https://adventofcode.com/2025/day/%d/input", day))
	if err != nil {
		return fmt.Errorf("download AOC URL: %w", err)
	}

	err = os.WriteFile(fmt.Sprintf("inputs/%s.txt", filename), []byte(strings.TrimSuffix(string(input), "\n")), 0664)
	if err != nil {
		return fmt.Errorf("write inputs/%s.txt: %w", filename, err)
	}

	return nil
}

func getData(filename string, lineHandler func(line string)) string {
	var err error
	stdinStat, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	var file io.ReadCloser
	if (stdinStat.Mode()&os.ModeCharDevice) == 0 && stdinStat.Size() > 0 {
		file = os.Stdin
	} else {
		file, err = inputs.Sets.Open(fmt.Sprintf("%s.txt", filename))
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("input %s not found, attempting to download...\n", filename)

			err = downloadInput(filename)
			if err != nil {
				panic(err)
			}

			fmt.Println("successfully downloaded input", filename)
			file, err = os.Open(fmt.Sprintf("inputs/%s.txt", filename))
		}

		if err != nil {
			panic(err)
		}
	}

	defer file.Close()

	if lineHandler != nil {
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			lineHandler(scanner.Text())
		}

		return ""
	}

	contents, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	return string(contents)
}

func GetStringContents(filename string) string {
	var retval string
	getData(filename, func(line string) {
		if len(retval) != 0 {
			panic("tried to parse multi-line file as a single line")
		}
		retval = line
	})
	return retval
}

func GetString(filename string) string {
	data := getData(filename, nil)
	return data
}

func GetStringLines(filename string) []string {
	retval := make([]string, 0)
	getData(filename, func(line string) {
		retval = append(retval, line)
	})
	return retval
}

func GetIntLines(filename string) []int64 {
	retval := make([]int64, 0)
	getData(filename, func(line string) {
		val, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			panic(err)
		}
		retval = append(retval, val)
	})
	return retval
}
