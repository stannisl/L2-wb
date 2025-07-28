package sort

import (
	"L2-10/args"
	"sort"
	"strconv"
	"strings"
)

type sortableLines struct {
	lines   []string
	flags   args.Flags
	columns [][]string
}

func newSortableLines(lines []string, flags args.Flags) *sortableLines {
	s := &sortableLines{
		lines: lines,
		flags: flags,
	}

	if flags.SortByColumn {
		s.columns = make([][]string, len(lines))
		for i, line := range lines {
			s.columns[i] = strings.Split(line, "\t")
		}
	}

	return s
}

func (s *sortableLines) Len() int {
	return len(s.lines)
}

func (s *sortableLines) Swap(i, j int) {
	s.lines[i], s.lines[j] = s.lines[j], s.lines[i]
	if s.columns != nil {
		s.columns[i], s.columns[j] = s.columns[j], s.columns[i]
	}
}

func (s *sortableLines) Less(i, j int) bool {
	key1 := s.getKey(i)
	key2 := s.getKey(j)

	var less bool
	if s.flags.SortByNumber {
		num1, err1 := strconv.ParseFloat(key1, 64)
		num2, err2 := strconv.ParseFloat(key2, 64)
		if err1 != nil {
			num1 = 0
		}
		if err2 != nil {
			num2 = 0
		}
		if num1 == num2 {
			less = s.lines[i] < s.lines[j]
		} else {
			less = num1 < num2
		}
	} else {
		if key1 == key2 {
			less = s.lines[i] < s.lines[j]
		} else {
			less = key1 < key2
		}
	}

	if s.flags.ReverseSort {
		return !less
	}
	return less
}

func (s *sortableLines) getKey(i int) string {
	if s.flags.SortByColumn {
		colIndex := s.flags.Column - 1
		if colIndex < 0 {
			colIndex = 0
		}
		if colIndex < len(s.columns[i]) {
			return s.columns[i][colIndex]
		}
		return ""
	}
	return s.lines[i]
}

// Lines сортирует строки в соответствии с указанными флагами.
func Lines(lines []string, flags args.Flags) []string {
	if len(lines) <= 1 {
		return lines
	}

	s := newSortableLines(lines, flags)
	sort.Sort(s)

	if flags.UniqueSort {
		return removeDuplicates(s.lines)
	}
	return s.lines
}

func removeDuplicates(lines []string) []string {
	if len(lines) <= 1 {
		return lines
	}

	i := 0
	for j := 1; j < len(lines); j++ {
		if lines[i] != lines[j] {
			i++
			lines[i] = lines[j]
		}
	}
	return lines[:i+1]
}
