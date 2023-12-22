package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scn := bufio.NewScanner(os.Stdin)
	var stacks1, stacks2 []Stack

	// parse the stack
	for {
		scn.Scan()
		txt := scn.Text()
		if txt == "" {
			break
		}

		for i, j := 0, 0; i < len(txt); i, j = i+4, j+1 {
			if stacks1 == nil && stacks2 == nil {
				stacks1 = make([]Stack, (len(txt)+1)/4)
				stacks2 = make([]Stack, (len(txt)+1)/4)
			}

			if txt[i] != '[' {
				continue
			}
			stacks1[j].PushBack(txt[i+1])
			stacks2[j].PushBack(txt[i+1])
		}
	}

	// parse the moves
	var moves []Move
	for scn.Scan() {
		txt := scn.Text()
		var n, src, dest int
		fmt.Sscanf(txt, "move %d from %d to %d", &n, &src, &dest)
		moves = append(moves, Move{
			n:    n,
			src:  src - 1,
			dest: dest - 1,
		})
	}

	// Run the simulation
	stacks1 = RunCrateMover9000(moves, stacks1)
	stacks2 = RunCrateMover9001(moves, stacks2)

	tops1 := ""
	tops2 := ""
	for _, s := range stacks1 {
		tops1 += string(s.Top())
	}

	for _, s := range stacks2 {
		tops2 += string(s.Top())
	}
	fmt.Println("part1: ", tops1)
	fmt.Println("part2: ", tops2)
}

func RunCrateMover9000(moves []Move, ss []Stack) []Stack {
	for _, m := range moves {
		for i := 0; i < m.n; i++ {
			ss[m.dest].Push(ss[m.src].Pop())
		}
	}
	return ss
}

func RunCrateMover9001(moves []Move, ss []Stack) []Stack {
	for _, m := range moves {
		tmpStack := Stack{}
		for i := 0; i < m.n; i++ {
			tmpStack.Push(ss[m.src].Pop())
		}
		for tmpStack.Size > 0 {
			ss[m.dest].Push(tmpStack.Pop())
		}
	}
	return ss
}

type Move struct {
	n    int
	src  int
	dest int
}
