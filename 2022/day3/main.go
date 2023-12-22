package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	b, _ := io.ReadAll(os.Stdin)
	fmt.Println("part 1:", part1(b))
	fmt.Println("part 2:", part2(b))
}

func part1(input []byte) int {
	getDupChar := func(a, b string) rune {
		counter := map[rune]struct{}{}
		for _, x := range a {
			counter[x] = struct{}{}
		}
		for _, y := range b {
			if _, ok := counter[y]; ok {
				return y
			}
		}
		return 0
	}

	res := 0
	scn := bufio.NewScanner(bytes.NewBuffer(input))
	for scn.Scan() {
		str := scn.Text()
		strLen := len(str)
		compartmentA := str[:strLen/2]
		compartmentB := str[strLen/2:]
		dupChar := getDupChar(compartmentA, compartmentB)
		res += getCharValue(dupChar)
	}
	return res
}

func part2(input []byte) int {
	getDupChar := func(a, b, c string) rune {
		counter := map[rune]struct{}{}
		var res rune
		for _, x := range a {
			counter[x] = struct{}{}
		}
		candidates := map[rune]struct{}{}
		for _, y := range b {
			if _, ok := counter[y]; ok {
				candidates[y] = struct{}{}
			}
		}
		for _, z := range c {
			if _, ok := candidates[z]; ok {
				res = z
			}
		}
		return res
	}

	res := 0
	scn := bufio.NewScanner(bytes.NewBuffer(input))

	var rucksacks [3]string
	for i := 0; scn.Scan(); i++ {
		str := scn.Text()
		rucksacks[i%3] = str
		if i%3 == 2 {
			fmt.Println(rucksacks)
			dupChar := getDupChar(rucksacks[0], rucksacks[1], rucksacks[2])
			fmt.Println(string(dupChar))
			res += getCharValue(dupChar)
		}
	}
	return res
}

func getCharValue(r rune) int {
	priority := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i, c := range priority {
		if c == r {
			return i + 1
		}
	}
	return 0
}
