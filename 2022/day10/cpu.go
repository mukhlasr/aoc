package main

import (
	"log"
	"strconv"
)

type simpleCPU struct {
	X          int                // Register
	Cycle      int                // Cycle
	Interrupts []func(*simpleCPU) // interrupts to be run every cycle
}

func (cpu *simpleCPU) do(c cmd, params ...string) {
	switch c {
	case noop:
		if len(params) != 0 {
			log.Fatalln("bad instruction, stopping")
		}
		cpu.increaseCycle()
		return
	case addx:
		if len(params) != 1 {
			log.Fatalln("bad instruction, stopping")
		}
		cpu.increaseCycle()
		cpu.increaseCycle()
		v, _ := strconv.Atoi(params[0])
		cpu.X += v
	}
}

func (cpu *simpleCPU) registerInterrupt(f func(*simpleCPU)) {
	if cpu.Interrupts == nil {
		cpu.Interrupts = []func(*simpleCPU){}
	}
	cpu.Interrupts = append(cpu.Interrupts, f)
}

func (cpu *simpleCPU) increaseCycle() {
	for _, f := range cpu.Interrupts {
		f(cpu)
	}
	cpu.Cycle++
}

type cmd string

const (
	noop cmd = "noop"
	addx cmd = "addx"
)
