package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, _ := io.ReadAll(os.Stdin)
	g1 := parseGame(bytes.NewBuffer(input), 3)
	g2 := parseGame(bytes.NewBuffer(input), 1)

	g1.play(20)
	fmt.Println("part 1:", g1.calculateMonkeyBusinessLevel())

	g2.play(10000)
	fmt.Println("part 2:", g2.calculateMonkeyBusinessLevel())
}

func parseGame(file io.Reader, worryLevelDividedBy int) *Game {
	scn := bufio.NewScanner(file)
	g := &Game{worryLevelDividedBy: worryLevelDividedBy}
	for scn.Scan() {
		text := scn.Text()
		if strings.HasPrefix(text, "Monkey") {
			m := monkey{}

			scn.Scan()
			m.items = parseStartingItems(scn.Text())

			scn.Scan()
			m.operation = parseOperations(scn.Text())

			testStr := ""
			for i := 0; i < 3; i++ {
				scn.Scan()
				testStr += fmt.Sprintln(scn.Text())
			}
			m.test = parseTestCondition(testStr)
			g.addMonkey(m)
		}
	}
	return g
}

type Game struct {
	monkeys             []*monkey
	worryLevelDividedBy int
}

func (g *Game) play(nRound int) {
	if nRound < 1 {
		return
	}

	for i := 0; i < nRound; i++ {
		for _, m := range g.monkeys {
			for len(m.items) > 0 {
				itemWorryLevel := m.inspect()
				if g.worryLevelDividedBy != 1 {
					itemWorryLevel /= g.worryLevelDividedBy
				} else {
					// when the worryLevel is not reduced by division, or in this case divided by 1
					// We can use this properties in the modular arithmetic
					// `(x mod kn) mod n == x mod n`
					// to reduce the number since we are only care about
					// its modulo result.
					itemWorryLevel %= g.productOfAllDivider()
				}
				target := m.test.throwTargetIdx(itemWorryLevel)
				g.monkeys[target].receive(itemWorryLevel)
			}
		}
	}
}

func (g *Game) addMonkey(m monkey) {
	g.monkeys = append(g.monkeys, &m)
}

func (g *Game) calculateMonkeyBusinessLevel() int {
	res := []int{}
	for _, m := range g.monkeys {
		res = append(res, m.inspected)
	}
	if len(res) == 0 {
		return 0
	}
	sort.Sort(sort.Reverse(sort.IntSlice(res)))
	return res[0] * res[1]
}

func (g *Game) productOfAllDivider() int {
	res := 1
	for _, v := range g.monkeys {
		res *= v.test.DivBy
	}
	return res
}

type monkey struct {
	items     []int
	operation operation
	test      testCondition
	inspected int
}

func (m *monkey) inspect() int {
	m.inspected++
	old := m.items[0]
	m.items = m.items[1:]
	return m.operation.eval(old)
}

func (m *monkey) receive(item int) {
	m.items = append(m.items, item)
}

func (m *monkey) String() string {
	return fmt.Sprintf("%+v", *m)
}

type operation struct {
	val changeVal
	op  operator
}

func (o *operation) eval(old int) int {
	val := old
	if o.val != oldVal {
		val = o.val.value()
	}
	switch o.op {
	case mul:
		return old * val
	case add:
		return old + val
	}
	return old
}

type changeVal string

const (
	oldVal = "old"
)

func (v changeVal) value() int {
	if v == oldVal {
		return 0
	}
	val, err := strconv.Atoi(string(v))
	if err != nil {
		log.Fatalln(err)
	}
	return val
}

type operator string

const (
	mul operator = "*"
	add operator = "+"
)

type testCondition struct {
	DivBy    int
	IdxTrue  int
	IdxFalse int
}

func (c testCondition) throwTargetIdx(worryLevel int) int {
	if worryLevel%c.DivBy == 0 {
		return c.IdxTrue
	}
	return c.IdxFalse
}

func parseStartingItems(str string) []int {
	var res []int
	re := regexp.MustCompile("[0-9]+")
	resStrs := re.FindAllString(str, -1)
	for _, r := range resStrs {
		n, _ := strconv.Atoi(r)
		res = append(res, n)
	}
	return res
}

func parseOperations(str string) operation {
	re := regexp.MustCompile("old.*")
	strs := strings.Split(re.FindString(str), " ")
	op := operator(strs[1])
	changeValue := changeVal(strs[2])
	return operation{
		val: changeValue,
		op:  op,
	}
}

func parseTestCondition(str string) testCondition {
	re := regexp.MustCompile("[0-9]+$")
	lines := strings.Split(str, "\n")

	var divBy int
	var idxTrue, idxFalse int
	divBy, _ = strconv.Atoi(re.FindString(lines[0]))
	idxTrue, _ = strconv.Atoi(re.FindString(lines[1]))
	idxFalse, _ = strconv.Atoi(re.FindString(lines[2]))

	return testCondition{
		DivBy:    divBy,
		IdxTrue:  idxTrue,
		IdxFalse: idxFalse,
	}
}
