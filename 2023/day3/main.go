package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f := os.Stdin
	defer f.Close()

	s := GenerateSchematicFromFile(f)
	fmt.Println(s.SumPartNumbers())
	fmt.Println(s.SumGearRatios())
}

func GenerateSchematicFromFile(f io.Reader) Schematic {
	var s Schematic
	s.SymbolsPositions = map[Pos]byte{}

	var rawData [][]byte

	buf := bufio.NewScanner(f)

	height := 0
	for i := 0; buf.Scan(); i++ {
		line := buf.Text()
		if i == 0 {
			s.Width = len(line)
		}

		rawData = append(rawData, []byte(line))

		for j := 0; j < len(line); j++ {
			c := line[j]
			if isDot(byte(c)) {
				continue
			}

			if isNumber(byte(c)) {
				length := 1
				for k := j + 1; k < len(line); k++ {
					if !isNumber(line[k]) {
						break
					}
					length++
				}

				s.NumberPositions = append(s.NumberPositions, Pos{
					X:      j,
					Y:      i,
					Length: length,
				})

				j = j + length - 1
				continue
			}

			s.SymbolsPositions[Pos{
				X:      j,
				Y:      i,
				Length: 1,
			}] = byte(c)
		}

		height = i + 1
	}

	s.RawData = rawData
	s.Height = height

	return s
}

func isDot(c byte) bool {
	return c == '.'
}

func isNumber(c byte) bool {
	return c >= '0' && c <= '9'
}

type Pos struct {
	X      int
	Y      int
	Length int
}

type Schematic struct {
	RawData          [][]byte
	Width            int
	Height           int
	SymbolsPositions map[Pos]byte
	NumberPositions  []Pos
}

func (s Schematic) SumGearRatios() int {
	gearsCandidates := map[Pos][]int{}

	for _, p := range s.NumberPositions {
		poss := s.CalculateAdjacentPositions(p)
		for _, pos := range poss {
			v, ok := s.SymbolsPositions[pos]
			if !ok || v != '*' {
				continue
			}

			num := s.NumberAtPos(p)
			gearsCandidates[pos] = append(gearsCandidates[pos], num)

			break
		}
	}

	var res int

	for _, nums := range gearsCandidates {
		if len(nums) != 2 {
			continue
		}

		res += nums[0] * nums[1]
	}

	return res
}

func (s Schematic) SumPartNumbers() int {
	var res int

	for _, p := range s.NumberPositions {
		poss := s.CalculateAdjacentPositions(p)
		for _, pos := range poss {
			_, ok := s.SymbolsPositions[pos]
			if !ok {
				continue
			}

			num := s.NumberAtPos(p)
			res += num
			break
		}
	}

	return res
}

func (s Schematic) CalculateAdjacentPositions(pos Pos) []Pos {
	p := pos
	var res []Pos

	start := Pos{X: p.X - 1, Y: p.Y, Length: 1}
	end := Pos{X: p.X + p.Length, Y: p.Y, Length: 1}

	res = append(res, start)
	res = append(res, end)

	for i := start.X; i <= end.X; i++ {
		res = append(res, Pos{X: i, Y: p.Y - 1, Length: 1})
	}

	for i := start.X; i <= end.X; i++ {
		res = append(res, Pos{X: i, Y: p.Y + 1, Length: 1})
	}

	return res
}

func (s Schematic) NumberAtPos(pos Pos) int {
	line := s.RawData[pos.Y]
	num := 0

	for i := pos.X; i < pos.X+pos.Length; i++ {
		n := int(line[i] - '0')
		if i == 0 {
			num = n
			continue
		}

		num *= 10
		num += n

	}

	return num
}
