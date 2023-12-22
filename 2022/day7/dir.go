package main

import "fmt"

type Dir struct {
	Name   string
	Parent *Dir
	Dirs   map[string]*Dir
	Files  map[string]*File
}

func (dir Dir) GetSize() int {
	total := 0
	for _, f := range dir.Files {
		total += f.Size
	}

	for _, d := range dir.Dirs {
		total += d.GetSize()
	}
	return total
}

func (dir Dir) PrintStructure() {
	dir.PrintStructureWithIndentation("")
}

func (dir Dir) PrintStructureWithIndentation(indentation string) {
	fmt.Println(indentation, "-", dir.Name, "(dir)")
	for _, d := range dir.Dirs {
		d.PrintStructureWithIndentation(indentation + "  ")
	}
	for _, f := range dir.Files {
		fmt.Println(indentation+"  ", "-", f.Name, fmt.Sprintf("(file, %d)", f.Size))
	}
}

type File struct {
	Name string
	Size int
}

func (f File) String() string {
	return fmt.Sprintf("%d %s", f.Size, f.Name)
}
