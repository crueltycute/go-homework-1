package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

type directoryContent []os.FileInfo

func (content directoryContent) Len() int { return len(content) }
func (content directoryContent) Less(i int, j int) bool { return content[i].Name() < content[j].Name() }
func (content directoryContent) Swap(i int, j int) {
	temp := content[i]
	content[i] = content[j]
	content[j] = temp
}

func makeTreeBranch(level int, last os.FileInfo, element os.FileInfo, mas *[]bool) string {
	var branch strings.Builder

	for i := 0; i < level; i++ {
		if (*mas)[i] == true {
			branch.WriteString( "│")
		}
		branch.WriteString("\t")
	}

	if element == last {
		branch.WriteString("└───")
		(*mas)[level] = false
	} else {
		branch.WriteString("├───")
	}

	branch.WriteString(element.Name())

	if element.IsDir() {
		branch.WriteString("\n")
	}

	return branch.String()
}

func fileSize(element os.FileInfo) string {
	if element.Size() == 0 {
		return " (empty)\n"
	} else {
		return " (" + strconv.FormatInt(element.Size(), 10) + "b)\n"
	}
}

func showDir(out io.Writer, path string, printFiles bool, level int, mas []bool) error {
	var dirContent directoryContent
	dirContent, err := ioutil.ReadDir(path)

	if err != nil {
		return errors.New("can't read dir")
	}

	sort.Sort(dirContent)
	if len(dirContent) != 0 {
		if !printFiles && !dirContent[len(dirContent)-1].IsDir() {
			dirContent = dirContent[:len(dirContent)-1]
		}
	}

	for _, element := range dirContent {
		if element.IsDir() {
			if len(mas) <= level + 1 {
				mas = append(mas, false)
				mas[level] = true
			}

			fmt.Fprint(out, makeTreeBranch(level, dirContent[len(dirContent) - 1], element, &mas))

			if err := showDir(out, path + "/" + element.Name(), printFiles, level + 1, mas); err != nil {
				return err
			}
		} else if !element.IsDir() && printFiles {
			fmt.Fprint(out, makeTreeBranch(level, dirContent[len(dirContent) - 1], element, &mas), fileSize(element))
		}
	}

	mas[level] = false

	return nil
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	mas := make([]bool, 0, 1)
	mas = append(mas, true)
	return showDir(out, path, printFiles, 0, mas)
}