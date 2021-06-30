package d01

import (
	"bufio"
	"io"
	"strconv"
)

// Parse returns nil on failure.
func Parse(r io.Reader) []int {
	res := make([]int, 0)
	s := bufio.NewScanner(r)
	for s.Scan() {
		n, err := strconv.Atoi(s.Text())
		if err != nil {
			return nil
		}
		res = append(res, n)
	}
	if err := s.Err(); err != nil {
		return nil
	}
	return res
}

func Part1(nums []int) int {
	seen := make(map[int]struct{})
	for _, n := range nums {
		partner := 2020 - n
		if _, ok := seen[partner]; !ok {
			seen[n] = struct{}{}
		} else {
			return n * partner
		}
	}
	return 0
}
