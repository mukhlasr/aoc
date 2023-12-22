package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scn := bufio.NewScanner(os.Stdin)
	var forest [][]forestTree
	for scn.Scan() {
		text := scn.Text()
		var row []forestTree
		for _, x := range text {
			row = append(row, forestTree{
				height: x,
			})
		}
		forest = append(forest, row)
	}

	interpretForestTrees(forest)

	fmt.Println("part 1:", numOfVisibleTrees(forest))
	fmt.Println("part 2:", highestScenicScore(forest))
}

func interpretForestTrees(forest [][]forestTree) {
	if len(forest) < 1 {
		return
	}

	if len(forest[0]) < 1 {
		return
	}
	forestRows := len(forest)
	forestCols := len(forest[0])

	for i := 0; i < forestRows; i++ {
		for j := 0; j < forestCols; j++ {
			t := forest[i][j]
			forest[i][j].visiblities.n = true
			forest[i][j].visiblities.s = true
			forest[i][j].visiblities.w = true
			forest[i][j].visiblities.e = true

			if i == 0 || j == 0 || i == forestRows-1 || j == forestCols-1 {
				continue
			}

			// west
			for x := j - 1; x >= 0; x-- {
				forest[i][j].nVisibleTrees.w++
				if forest[i][x].height >= t.height {
					forest[i][j].visiblities.w = false
					break
				}
			}

			// east
			for x := j + 1; x < forestCols; x++ {
				forest[i][j].nVisibleTrees.e++
				if forest[i][x].height >= t.height {
					forest[i][j].visiblities.e = false
					break
				}
			}

			// north
			for x := i - 1; x >= 0; x-- {
				forest[i][j].nVisibleTrees.n++
				if forest[x][j].height >= t.height {
					forest[i][j].visiblities.n = false
					break
				}
			}

			// south
			for x := i + 1; x < forestRows; x++ {
				forest[i][j].nVisibleTrees.s++
				if forest[x][j].height >= t.height {
					forest[i][j].visiblities.s = false
					break
				}
			}
		}
	}
}

func numOfVisibleTrees(forest [][]forestTree) int {
	res := 0
	if len(forest) < 1 {
		return res
	}

	if len(forest[0]) < 1 {
		return res
	}

	for i := 0; i < len(forest); i++ {
		for j := 0; j < len(forest[0]); j++ {
			t := forest[i][j]
			if t.isVisible() {
				res++
			}
		}
	}

	return res
}

func highestScenicScore(forest [][]forestTree) int {
	res := 0
	for i := 0; i < len(forest); i++ {
		for j := 0; j < len(forest[0]); j++ {
			t := forest[i][j]
			if t.scenicScore() > res {
				res = t.scenicScore()
			}
		}
	}
	return res
}

type forestTree struct {
	height      rune
	visiblities struct {
		n bool
		s bool
		w bool
		e bool
	}
	nVisibleTrees struct {
		n int
		s int
		w int
		e int
	}
}

func (t forestTree) isVisible() bool {
	return t.visiblities.n || t.visiblities.s || t.visiblities.w || t.visiblities.e
}

func (t forestTree) scenicScore() int {
	return t.nVisibleTrees.n * t.nVisibleTrees.s * t.nVisibleTrees.w * t.nVisibleTrees.e
}
