package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scn := bufio.NewScanner(os.Stdin)
	cpu := &simpleCPU{X: 1, Cycle: 1}
	var result []int

	sampleValueAtCertainCycles := func(cpu *simpleCPU) {
		for _, c := range []int{20, 60, 100, 140, 180, 220} {
			if c == cpu.Cycle {
				result = append(result, cpu.Cycle*cpu.X)
			}
		}
	}

	drawCRT := func(cpu *simpleCPU) {
		crtPosition := (cpu.Cycle - 1) % 40
		pixelValue := "."
		if crtPosition == cpu.X || crtPosition == cpu.X+1 || crtPosition == cpu.X-1 {
			pixelValue = "#"
		}
		fmt.Print(pixelValue)

		if cpu.Cycle%40 == 0 { // write new row
			fmt.Println()
		}
	}

	cpu.registerInterrupt(sampleValueAtCertainCycles)
	cpu.registerInterrupt(drawCRT)

	for scn.Scan() {
		text := strings.Split(scn.Text(), " ")
		cmd := cmd(text[0])
		params := text[1:]
		cpu.do(cmd, params...)
	}
	sum := 0
	for _, x := range result {
		sum += x
	}
	fmt.Println("part1:", sum)
}
