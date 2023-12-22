package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f := os.Stdin
	defer f.Close()

	buf := bufio.NewScanner(f)

	res1 := 0
	res2 := 0

	for buf.Scan() {
		line := buf.Text()
		rec := getGameRecord(line)
		if isValidGame(rec.CubeRevealedStr) {
			res1 += rec.ID
		}

		res2 += sumOfThePower(rec.CubeRevealedStr)
	}

	fmt.Println(res1)
	fmt.Println(res2)
}

type GameRecord struct {
	ID              int
	CubeRevealedStr string
}

func getGameRecord(line string) GameRecord {
	splitted := strings.Split(line, ":")

	prefix := splitted[0]
	cubeRevealedStr := splitted[1]

	idStr := strings.Split(prefix, " ")[1]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}

	return GameRecord{
		ID:              id,
		CubeRevealedStr: cubeRevealedStr,
	}
}

func isValidGame(cubeRevealedStr string) bool {
	maxCubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	for _, str := range strings.Split(cubeRevealedStr, ";") {
		for _, cube := range strings.Split(str, ",") {
			cube := strings.Trim(cube, " ")
			splitted := strings.Split(cube, " ")
			color := splitted[1]
			number, err := strconv.Atoi(splitted[0])
			if err != nil {
				panic(err)
			}

			if number > maxCubes[color] {
				return false
			}
		}
	}

	return true
}

func sumOfThePower(cubeRevealedStr string) int {
	maxCubes := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for _, str := range strings.Split(cubeRevealedStr, ";") {
		for _, cube := range strings.Split(str, ",") {
			cube := strings.Trim(cube, " ")
			splitted := strings.Split(cube, " ")
			color := splitted[1]
			number, err := strconv.Atoi(splitted[0])
			if err != nil {
				panic(err)
			}

			if number > maxCubes[color] {
				maxCubes[color] = number
			}
		}
	}

	res := 1
	for _, v := range maxCubes {
		res *= v
	}

	return res
}
