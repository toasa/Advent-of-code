package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Dir struct {
	name   string
	parent *Dir

	files map[string]*File
	dirs  map[string]*Dir

	size uint64
}

type File struct {
	name string
	size uint64
}

func newDir(name string, parent *Dir) *Dir {
	return &Dir{
		name:   name,
		parent: parent,
		files:  make(map[string]*File),
		dirs:   make(map[string]*Dir),
	}
}

func newFile(name string, size uint64) *File {
	return &File{
		name: name,
		size: size,
	}
}

func splitIntoLines(filename string) []string {
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
	CMD_CD = iota
	CMD_LS
)

type parseResult struct {
	cmd int

	// For cd command
	dstDir string

	// For ls command
	dirs  []string
	files []*File
}

func parse(lines []string) (parseResult, int) {
	if lines[0][0] != '$' {
		panic("Cmd not found")
	}

	p := parseResult{}
	i := 0

	op := lines[0][2:4]
	switch op {
	case "cd":
		p.dstDir = lines[0][len("$ cd "):]
		p.cmd = CMD_CD
		i = 1
	case "ls":
		p.cmd = CMD_LS

		lines = lines[1:]
		j := 0
		for ; j < len(lines); j++ {
			// Found next command
			if lines[j][0] == '$' {
				break
			}

			if lines[j][:3] == "dir" {
				// Directory
				p.dirs = append(p.dirs, lines[j][4:])
			} else {
				// File
				sp := strings.Index(lines[j], " ")
				name := lines[j][sp+1:]
				size, _ := strconv.Atoi(lines[j][:sp])

				f := newFile(name, uint64(size))
				p.files = append(p.files, f)
			}
		}

		i = j + 1

	default:
		panic("Invalid op")
	}

	return p, i
}

func buildFileTree(lines []string) *Dir {
	// skip root dir
	lines = lines[1:]

	root := newDir("/", nil)
	curDir := root

	for len(lines) > 0 {
		res, i := parse(lines)
		lines = lines[i:]

		// fmt.Println(res)
		switch res.cmd {
		case CMD_CD:
			if res.dstDir == ".." {
				curDir = curDir.parent
			} else {
				d, ok := curDir.dirs[res.dstDir]
				if !ok {
					panic("CD Dir not found")
				}
				curDir = d
			}
		case CMD_LS:
			for _, dirName := range res.dirs {
				if _, ok := curDir.dirs[dirName]; !ok {
					// Create child directory
					d := newDir(dirName, curDir)
					curDir.dirs[dirName] = d
				}
			}

			for _, f := range res.files {
				if _, ok := curDir.files[f.name]; !ok {
					curDir.files[f.name] = f
				}
			}
		}

	}

	return root
}

func calcDirSize(dir *Dir) {
	var size uint64 = 0
	for _, d := range dir.dirs {
		calcDirSize(d)
		size += d.size
	}
	for _, f := range dir.files {
		size += f.size
	}

	dir.size = size
}

const FILE_SIZE_LIMIT = 100000

func _solvePart1(dir *Dir) uint64 {
	var res uint64 = 0
	if dir.size <= FILE_SIZE_LIMIT {
		res += dir.size
	}

	for _, d := range dir.dirs {
		res += _solvePart1(d)
	}
	return res

}

func solvePart1(lines []string) {
	root := buildFileTree(lines)
	calcDirSize(root)

	fmt.Println(_solvePart1(root))
}

func main() {
	filename := "input"
	lines := splitIntoLines(filename)

	solvePart1(lines)
}
