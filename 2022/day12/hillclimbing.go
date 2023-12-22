package main

import (
	"fmt"
	"math"
	"sort"
)

func hillClimbing(hill [][]byte, s, e position) []position {
	fmt.Println(len(hill), len(hill[0]))
	isAllowed := func(curChar, targetChar byte) bool {
		return characterValue(curChar)+1 == characterValue(targetChar) ||
			characterValue(curChar) >= characterValue(targetChar)
	}

	euclideanDistance := func(a, b position) float64 {
		x2 := (a.x - b.x) * (a.x - b.x)
		y2 := (a.y - b.y) * (a.y - b.y)
		return math.Abs(math.Sqrt(float64(x2) + float64(y2)))
	}

	visited := map[position]bool{}
	visited[s] = true

	var st Stack
	prevMap := map[position]position{}

	var search func(position, position)
	search = func(s, e position) {
		visited[s] = true

		if s == e {
			return
		}

		var candidates []position
		for _, pos := range []position{
			{x: s.x, y: s.y - 1},
			{x: s.x, y: s.y + 1},
			{x: s.x - 1, y: s.y},
			{x: s.x + 1, y: s.y}} {
			if pos.x < 0 || pos.y < 0 || pos.y >= len(hill) || pos.x >= len(hill[0]) {
				continue
			}
			if visited[pos] || !isAllowed(hill[s.y][s.x], hill[pos.y][pos.x]) {
				continue
			}
			candidates = append(candidates, pos)
		}

		sort.Slice(candidates, func(i, j int) bool {
			pos1 := candidates[i]
			pos2 := candidates[j]
			return euclideanDistance(pos1, e) > euclideanDistance(pos2, e)
		})

		for _, c := range candidates {
			st.Push(c)
		}

		if st.Size == 0 {
			return
		}

		next := st.Pop()

		prevMap[next] = s
		search(next, e)
	}

	search(s, e)

	fmt.Println(len(prevMap))
	var res []position
	count := 0
	for prev, ok := prevMap[e]; ok; prev, ok = prevMap[prev] {
		count++
		fmt.Println(count)
	}
	return res
}

func characterValue(b byte) int {
	if b == 'S' {
		b = 'a'
	}
	if b == 'E' {
		b = 'z'
	}

	charVal := b - 'a' + 1 // a = 1, b = 2, ..., z = 26
	return int(charVal)
}

type position struct {
	x int
	y int
}

type path struct {
	pos  position
	next *path
}

func (p *path) length() int {
	count := 0
	pp := &p
	for (*pp) != nil {
		count++
		pp = &(*pp).next
	}
	return count
}
