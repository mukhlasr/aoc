package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f := os.Stdin

	buf := bufio.NewScanner(f)

	var cards []Card
	for buf.Scan() {
		cards = append(cards, ParseCard(buf.Text()))
	}

	var res int
	for _, c := range cards {
		res += c.CalculatePoint()
	}

	fmt.Println(res)
	fmt.Println(CalculatedTotalCardsWithDuplicates(cards))
}

func ParseCard(line string) Card {
	line = removeDuplicatedSpace(line)
	var c Card
	split := strings.Split(line, ":")
	numsStr := split[1]
	prefix := split[0]

	gameIDStr := strings.Split(prefix, " ")[1]

	id, err := strconv.Atoi(gameIDStr)
	if err != nil {
		log.Fatalln(err)
	}

	numsSplited := strings.Split(numsStr, "|")
	winningNumsStr := strings.Trim(numsSplited[0], " ")
	numsStr = strings.Trim(numsSplited[1], " ")

	var winningNums []int

	for _, n := range strings.Split(winningNumsStr, " ") {
		n := strings.Trim(n, " ")

		num, err := strconv.Atoi(n)
		if err != nil {
			log.Fatalln(err)
		}

		winningNums = append(winningNums, num)
	}

	var nums []int

	for _, n := range strings.Split(numsStr, " ") {
		n := strings.Trim(n, " ")
		num, err := strconv.Atoi(n)
		if err != nil {
			log.Fatalln(err)
		}

		nums = append(nums, num)
	}

	c.ID = id
	c.WinningNums = winningNums
	c.Nums = nums
	c.WinningNumbersAppeared = c.FindWinningNumbersApppeared()

	return c
}

func removeDuplicatedSpace(str string) string {
	var res []byte
	before := byte(' ')
	for i := 0; i < len(str); i++ {
		if str[i] == before && str[i] == ' ' {
			continue
		}

		res = append(res, str[i])

		before = str[i]
	}

	return string(res)
}

type Card struct {
	ID                     int
	WinningNums            []int
	Nums                   []int
	WinningNumbersAppeared []int
}

func (c Card) FindWinningNumbersApppeared() []int {
	numMap := map[int]struct{}{}
	for _, n := range c.WinningNums {
		numMap[n] = struct{}{}
	}

	var res []int
	for _, n := range c.Nums {
		if _, ok := numMap[n]; ok {
			res = append(res, n)
		}
	}

	return res
}

func (c Card) CalculatePoint() int {
	res := 1
	n := len(c.WinningNumbersAppeared)
	if n == 0 {
		return 0
	}

	return res << (n - 1)
}

func (c Card) GetDuplicateCardsIDs() []int {
	var res []int
	n := len(c.WinningNumbersAppeared)
	for i := 0; i < n; i++ {
		res = append(res, c.ID+i+1)
	}

	return res

}

func CalculatedTotalCardsWithDuplicates(cards []Card) int {
	mapIDToCard := map[int]Card{}
	for _, c := range cards {
		mapIDToCard[c.ID] = c
	}

	nDuplicates := 0

	for _, c := range cards {
		var dupCardIDs []int
		// append duplicates
		dupCardIDs = append(dupCardIDs, c.GetDuplicateCardsIDs()...)

		for i := 0; i < len(dupCardIDs); i++ { // find duplicates of duplicates
			d := dupCardIDs[i]
			c := mapIDToCard[d]
			dupCardIDs = append(dupCardIDs, c.GetDuplicateCardsIDs()...)
		}

		nDuplicates += len(dupCardIDs)
	}

	// number of cards + number of duplicates = total cards
	return nDuplicates + len(cards)
}
