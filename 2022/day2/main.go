package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type move int
type matchResult int

const (
	ROCK move = iota
	PAPER
	SCISSOR

	WIN matchResult = iota
	LOSE
	DRAW
)

var ()

func main() {
	scn := bufio.NewScanner(os.Stdin)
	totalScore1 := 0
	totalScore2 := 0
	for scn.Scan() {
		moves := strings.Split(scn.Text(), " ")
		totalScore1 += part1(moves[0], moves[1])
		totalScore2 += part2(moves[0], moves[1])
	}
	fmt.Println("part1:", totalScore1)
	fmt.Println("part2:", totalScore2)
}

func part1(elfMoveStr, playerMoveStr string) int {
	elfMove := getElfMove(elfMoveStr)
	playerMove := map[string]move{
		"X": ROCK,
		"Y": PAPER,
		"Z": SCISSOR,
	}[playerMoveStr]
	res := play(elfMove, playerMove)
	matchScore := getMoveScore(playerMove) + getMatchScore(res)
	return matchScore
}

func part2(elfMoveStr, playerMoveStr string) int {
	elfMove := getElfMove(elfMoveStr)
	winningMove := map[move]move{
		ROCK:    PAPER,
		PAPER:   SCISSOR,
		SCISSOR: ROCK,
	}
	losingMove := map[move]move{
		ROCK:    SCISSOR,
		SCISSOR: PAPER,
		PAPER:   ROCK,
	}
	if playerMoveStr == "X" {
		move := losingMove[elfMove]
		return getMoveScore(move) + getMatchScore(LOSE)
	}
	if playerMoveStr == "Z" {
		move := winningMove[elfMove]
		return getMoveScore(move) + getMatchScore(WIN)
	}
	return getMoveScore(elfMove) + getMatchScore(DRAW)
}

func play(elf, player move) matchResult {
	if elf == player {
		return DRAW
	}
	if elf == ROCK && player == PAPER {
		return WIN
	}
	if elf == PAPER && player == SCISSOR {
		return WIN
	}
	if elf == SCISSOR && player == ROCK {
		return WIN
	}
	return LOSE

}

func getMatchScore(res matchResult) int {
	return map[matchResult]int{
		WIN:  6,
		DRAW: 3,
		LOSE: 0,
	}[res]

}

func getMoveScore(m move) int {
	return map[move]int{
		ROCK:    1,
		PAPER:   2,
		SCISSOR: 3,
	}[m]
}

func getElfMove(str string) move {
	return map[string]move{
		"A": ROCK,
		"B": PAPER,
		"C": SCISSOR,
	}[str]
}
