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
	scn := bufio.NewScanner(os.Stdin)
	var d *Dir
	lastCommand := ""
	cmdResult := ""

	// begin parsing
	for scn.Scan() {
		text := scn.Text()
		if lastCommand == "ls" {
			for !isCommand(text) {
				cmdResult += text + "\n"
				if !scn.Scan() {
					break
				}
				text = scn.Text()
			}
			d.Dirs, d.Files = parseLsResult(d, cmdResult)
			cmdResult = ""
		}

		cmd, _ := ParseCommand(text)
		switch lastCommand = cmd.Name; lastCommand {
		case "cd":
			switch dirName := cmd.Params[0]; dirName {
			case "":
				log.Fatalln("invalid dir")
			case "..":
				d = d.Parent
			default:
				if d == nil {
					d = &Dir{
						Name: dirName,
					}
					continue
				}
				if target, ok := d.Dirs[dirName]; ok {
					d = target
					continue
				}
				log.Fatalln("invalid cd: dir " + dirName + " not found")

			}
		case "ls":
		}
	}

	// back to the root dir
	for d.Parent != nil {
		d = d.Parent
	}

	// d.PrintStructure()
	fmt.Println("part 1:", totalDirSizeBelowThreshold(d, 100000))

	freeSize := 70000000 - d.GetSize()
	fmt.Println("part 2:", smallestDirSizeAboveThreshold(d, 30000000-freeSize))
}

func totalDirSizeBelowThreshold(dir *Dir, threshold int) int {
	total := 0
	var calculate func(dir *Dir)
	calculate = func(dir *Dir) {
		s := dir.GetSize()
		if s <= threshold {
			total += s
		}

		for _, v := range dir.Dirs {
			calculate(v)
		}
	}

	calculate(dir)
	return total
}

func smallestDirSizeAboveThreshold(dir *Dir, threshold int) int {
	res := dir.GetSize()
	var calculate func(dir *Dir)
	calculate = func(dir *Dir) {
		s := dir.GetSize()
		if s < threshold {
			return
		}

		if s < res {
			res = s
		}

		for _, v := range dir.Dirs {
			calculate(v)
		}
	}
	calculate(dir)
	return res
}

func parseLsResult(curDir *Dir, text string) (map[string]*Dir, map[string]*File) {
	const DirPrefix = "dir "
	scn := bufio.NewScanner(strings.NewReader(text))
	dirResult := make(map[string]*Dir)
	fileResult := make(map[string]*File)
	for scn.Scan() {
		text := scn.Text()
		if strings.HasPrefix(text, DirPrefix) {
			dirName := text[len(DirPrefix):]
			dirResult[dirName] = &Dir{
				Name:   dirName,
				Parent: curDir,
			}
			continue
		}
		f := strings.Split(text, " ")
		filename := f[1]
		filesize, _ := strconv.Atoi(f[0])
		fileResult[filename] = &File{
			Name: filename,
			Size: filesize,
		}
	}
	return dirResult, fileResult
}
