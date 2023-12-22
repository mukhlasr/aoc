package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f := os.Stdin
	defer f.Close()

	buf := bufio.NewScanner(f)

	res1 := 0
	res2 := 0

	for buf.Scan() {
		line := buf.Text()
		res1 += findNum(line)
		res2 += findNumPart2(line)
	}

	fmt.Println(res1)
	fmt.Println(res2)
}

func findNum(line string) int {
	res := 0
	for i := 0; i < len(line); i++ {
		if !isNumber(line[i]) {
			continue
		}
		res = int(line[i] - '0')
		break
	}

	for j := len(line) - 1; j >= 0; j-- {
		if !isNumber(line[j]) {
			continue
		}

		res = res*10 + int(line[j]-'0')
		break
	}

	return res
}

func findNumPart2(line string) int {
	nums := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	res := 0
outer:
	for i := 0; i < len(line); i++ {
		if isNumber(line[i]) {
			res = int(line[i] - '0')
			break
		}

		for j := 3; j <= 5; j++ {
			if i+j > len(line)-1 {
				continue
			}

			str := line[i : i+j]
			if val, ok := nums[str]; ok {
				res = val
				break outer
			}
		}
	}

outer1:
	for j := len(line) - 1; j >= 0; j-- {
		if isNumber(line[j]) {
			res = res*10 + int(line[j]-'0')
			break
		}

		for k := 2; k <= 4; k++ {
			if j < k {
				continue
			}

			str := line[j-k : j+1]
			if val, ok := nums[str]; ok {
				res = res*10 + val
				break outer1
			}
		}
	}

	return res
}

func isNumber(b byte) bool {
	return b >= '0' && b <= '9'
}
