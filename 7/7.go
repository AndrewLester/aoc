package seven

import (
	"aoc2022/input"
	"math"
	"path"
	"sort"
	"strconv"
	"strings"
)

type File struct {
	name     string
	size     int
	children []*File
}

const totalDiskSpace = 70000000
const neededDiskSpace = 30000000

func parseCommand(command string, wd string) (string, string) {
	arguments := strings.Split(command, " ")
	switch arguments[0] {
	case "cd":
		if arguments[1] == "/" {
			return "cd", "/"
		} else if arguments[1] == ".." {
			return "cd", path.Dir(wd)
		} else {
			return "cd", path.Join(wd, arguments[1])
		}
	case "ls":
		return "ls", ""
	default:
		panic("Bad path")
	}
}

func getFileFromPath(fileSystem *File, wd string) *File {
	if wd[len(wd)-1] == '/' {
		return fileSystem
	}

	if wd[0] == '/' {
		wd = wd[1:]
	}

	splitWd := strings.SplitN(wd, "/", 2)
	fileName := splitWd[0]
	newWd := ""
	if len(splitWd) > 1 {
		newWd = splitWd[1]
	}

	for _, file := range fileSystem.children {
		if file.name == fileName {
			if newWd == "" {
				return file
			}
			return getFileFromPath(file, newWd)
		}
	}
	return nil
}

func buildFilesystem(text string) *File {
	lines := strings.Split(text, "\n")

	fileSystem := &File{size: -1, name: "/"}
	wd := "/"
	cmd := ""
	for _, line := range lines {
		if len(line) == 0 {
			break
		}

		if line[0] == '$' {
			new_cmd, new_wd := parseCommand(line[2:], wd)
			if new_wd != "" {
				wd = new_wd
			}
			cmd = new_cmd
			continue
		}

		if cmd != "ls" {
			// should never happen...
			continue
		}

		folder := getFileFromPath(fileSystem, wd)

		splitLine := strings.Split(line, " ")
		size, err := strconv.Atoi(splitLine[0])
		name := splitLine[1]

		if err != nil {
			size = -1
		}

		file := &File{
			name:     name,
			size:     size,
			children: []*File{},
		}

		folder.children = append(folder.children, file)
	}

	return fileSystem
}

func findSmallDirectories(parent *File, smallDirSizes *[]int) int {
	if len(parent.children) == 0 {
		return 0
	}

	totalSize := 0
	for _, file := range parent.children {
		if file.size == -1 {
			totalSize += findSmallDirectories(file, smallDirSizes)
		} else {
			totalSize += file.size
		}
	}

	if totalSize <= 100000 {
		*smallDirSizes = append(*smallDirSizes, totalSize)
	}

	return totalSize
}

func Part1() int {
	text := input.ReadInput(7)

	fileSystem := buildFilesystem(text)

	smallDirSizes := []int{}
	findSmallDirectories(fileSystem, &smallDirSizes)

	totalSize := 0
	for _, size := range smallDirSizes {
		totalSize += size
	}
	return totalSize
}

func findMatchingDirectories(parent *File, matchingDirs *[]int, needs int) int {
	if len(parent.children) == 0 {
		return 0
	}

	totalSize := 0
	for _, file := range parent.children {
		if file.size == -1 {
			totalSize += findMatchingDirectories(file, matchingDirs, needs)
		} else {
			totalSize += file.size
		}
	}

	if totalSize >= needs {
		*matchingDirs = append(*matchingDirs, totalSize)
	}

	return totalSize
}

func Part2() int {
	text := input.ReadInput(7)

	fileSystem := buildFilesystem(text)
	unused := []int{}
	size := findMatchingDirectories(fileSystem, &unused, math.MaxInt32)
	unusedSpace := totalDiskSpace - size

	matchingDirs := []int{}
	findMatchingDirectories(fileSystem, &matchingDirs, neededDiskSpace-unusedSpace)

	sort.Ints(matchingDirs)
	return matchingDirs[0]
}
