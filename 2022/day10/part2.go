package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getInputFromFile(filename string) []string {
	lines := []string{}

	fp, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	s := bufio.NewScanner(fp)

	for s.Scan() {
		lines = append(lines, s.Text())
	}

	if err = s.Err(); err != nil {
		panic(err)
	}

	return lines
}

const (
	OP_ADDX = iota
	OP_NOOP
)

type Inst struct {
	op  int
	val int
}

func parse(input []string) []Inst {
	insts := make([]Inst, len(input))
	for i, line := range input {
		switch line[:4] {
		case "addx":
			val, _ := strconv.Atoi(line[5:])
			insts[i] = Inst{op: OP_ADDX, val: val}
		case "noop":
			insts[i] = Inst{op: OP_NOOP}
		default:
			panic("invalid instruction")
		}
	}
	return insts
}

type CPU struct {
	cycle int
	reg   int
}

func newCPU() *CPU {
	return &CPU{
		cycle: 0,
		reg:   1,
	}
}

var CPT [240]byte

func (cpu *CPU) draw() {
	isOverLap := func(i, spriteI int) bool {
		if i == spriteI-1 || i == spriteI || i == spriteI+1 {
			return true
		}
		return false
	}

	if isOverLap(cpu.cycle%40, cpu.reg) {
		CPT[cpu.cycle] = '#'
	} else {
		CPT[cpu.cycle] = '.'
	}
}

func (cpu *CPU) exec(inst Inst) {
	switch inst.op {
	case OP_ADDX:
		cpu.draw()
		cpu.cycle++
		cpu.draw()
		cpu.cycle++
		cpu.reg += inst.val
	case OP_NOOP:
		cpu.draw()
		cpu.cycle++
	}
}

func (cpu *CPU) run(insts []Inst) {
	for _, inst := range insts {
		cpu.exec(inst)
	}
}

func solvePart1(input []string) {
	insts := parse(input)
	cpu := newCPU()
	cpu.run(insts)

	for i := 0; i < 6; i++ {
		for j := 0; j < 40; j++ {
			fmt.Printf("%c", CPT[i*40+j])
		}
		fmt.Println()
	}
}

func main() {
	filename := "input"
	input := getInputFromFile(filename)

	solvePart1(input)
}
