package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type elf struct {
	calories      []int
	totalCalories int
}

func main() {
	topThree := TopThree{}

	scn := bufio.NewScanner(os.Stdin)
	e := elf{}
	isMore := true
	for isMore {
		isMore = scn.Scan()
		txt := scn.Text()
		if txt == "" || !isMore {
			e.totalCalories = sumArr(e.calories)
			topThree.Insert(e.totalCalories)
			e = elf{}
			continue
		}
		calorie, _ := strconv.Atoi(txt)
		e.calories = append(e.calories, calorie)
	}
	fmt.Println("largest:", topThree.First())
	fmt.Println("top three:", topThree.Elems)
	fmt.Println("sum of top three:", topThree.Sum())
}

func sumArr(arr []int) int {
	var res int
	for _, a := range arr {
		res += a
	}
	return res
}

type TopThree struct {
	Elems [3]int
}

func (q *TopThree) First() int {
	return q.Elems[0]
}

func (q *TopThree) Insert(num int) {
	shiftRight := func(from int) {
		for i := 2; i > from; i-- {
			q.Elems[i] = q.Elems[i-1]
		}
	}

	for i, e := range q.Elems {
		if num > e {
			shiftRight(i)
			q.Elems[i] = num
			break
		}
	}
}

func (q *TopThree) Sum() int {
	res := 0
	for _, e := range q.Elems {
		res += e
	}
	return res
}
