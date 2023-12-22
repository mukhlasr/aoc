package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scn := bufio.NewScanner(os.Stdin)
	fullyContained := 0
	overlapped := 0
	for scn.Scan() {
		ranges := strings.Split(scn.Text(), ",")
		smallerRange, biggerRange := NewRangeFromString(ranges[0]), NewRangeFromString(ranges[1])
		if smallerRange.Length > biggerRange.Length {
			smallerRange, biggerRange = biggerRange, smallerRange
		}

		if biggerRange.Contains(smallerRange) {
			fullyContained++
		}

		if biggerRange.Overlaps(smallerRange) {
			overlapped++
		}
	}
	fmt.Println("fully contained:", fullyContained)
	fmt.Println("overlapped:", overlapped)
}

func NewRangeFromString(str string) Range {
	components := strings.Split(str, "-")
	start, _ := strconv.Atoi(components[0])
	end, _ := strconv.Atoi(components[1])
	return Range{
		Start:  start,
		End:    end,
		Length: end - start + 1,
	}
}

type Range struct {
	Start  int
	End    int
	Length int
}

func (r Range) Contains(a Range) bool {
	return a.Start >= r.Start && a.End <= r.End
}

func (r Range) Overlaps(a Range) bool {
	headOverlap := r.Start <= a.End && r.End >= a.End
	tailOverlap := r.End >= a.Start && r.Start <= a.Start
	return headOverlap || tailOverlap
}
