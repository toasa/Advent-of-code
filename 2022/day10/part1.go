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

var checkCycles []int = []int{20, 60, 100, 140, 180, 220}
var res int = 0

func (cpu *CPU) check() {
	if len(checkCycles) > 0 && cpu.cycle >= checkCycles[0] {
		res += checkCycles[0] * cpu.reg
		checkCycles = checkCycles[1:]
	}
}

func (cpu *CPU) exec(inst Inst) {
	switch inst.op {
	case OP_ADDX:
		cpu.cycle++
		cpu.check()
		cpu.cycle++
		cpu.check()
		cpu.reg += inst.val
	case OP_NOOP:
		cpu.cycle++
		cpu.check()
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

	fmt.Println(cpu.cycle)
	fmt.Println(res)
}

func main() {
	filename := "input"
	input := getInputFromFile(filename)

	solvePart1(input)
}
