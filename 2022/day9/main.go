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

	head := newRope(10)
	tail := head.tail()
	for scn.Scan() {
		text := strings.Split(scn.Text(), " ")
		d := direction(text[0])
		n, _ := strconv.Atoi(text[1])
		for i := 0; i < n; i++ {
			head.move(d)
		}

	}
	fmt.Println("part 1:", len(head.next.visited))
	fmt.Println("part 2:", len(tail.visited))
}

func newRope(length int) *node {
	if length < 1 {
		return nil
	}
	head := &node{
		visited: map[position]struct{}{},
	}
	head.setCurrentPosition(position{})
	head.next = newRope(length - 1)
	return head
}

type node struct {
	pos     position
	next    *node
	visited map[position]struct{}
}

func (n *node) move(d direction) {
	newPos := n.pos.move(d)
	n.setCurrentPosition(newPos)

	target := newPos
	pn := &n
	for (*pn).next != nil {
		n := *pn
		if n.pos.isNeighbor(n.next.pos) {
			return
		}
		newPos := n.next.pos.getNewPositionNearTarget(target)
		n.next.setCurrentPosition(newPos)
		target = newPos
		pn = &n.next
	}
}

func (n *node) setCurrentPosition(p position) {
	n.visited[p] = struct{}{}
	n.pos = p
}

func (n *node) tail() *node {
	pTail := &n
	for (*pTail).next != nil {
		pTail = &(*pTail).next
	}
	return *pTail
}

func (n *node) String() string {
	var pos []string
	pn := &n
	for (*pn) != nil {
		pos = append(pos, (*pn).pos.String())
		pn = &(*pn).next
	}
	return strings.Join(pos, " -> ")
}

type position struct {
	x int
	y int
}

func (p position) move(d direction) position {
	res := p
	switch d {
	case left:
		res.x--
	case right:
		res.x++
	case up:
		res.y--
	case down:
		res.y++
	}
	return res
}

func (p position) isNeighbor(target position) bool {
	if p == target {
		return true
	}

	n := p.move(up)
	s := p.move(down)
	w := p.move(left)
	e := p.move(right)
	nw := n.move(left)
	ne := n.move(right)
	sw := s.move(left)
	se := s.move(right)

	for _, p := range []position{n, s, w, e, nw, ne, sw, se} {
		if target == p {
			return true
		}
	}
	return false
}

func (p position) getNewPositionNearTarget(target position) position {
	if p == target {
		return p
	}

	if target.x > p.x {
		p.x++
	}

	if target.x < p.x {
		p.x--
	}

	if target.y > p.y {
		p.y++
	}

	if target.y < p.y {
		p.y--
	}

	return p

}

func (p position) String() string {
	return fmt.Sprintf("(x: %d, y: %d)", p.x, p.y)
}

type direction string

const (
	left  direction = "L"
	right direction = "R"
	up    direction = "U"
	down  direction = "D"
)
