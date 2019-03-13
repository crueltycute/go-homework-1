package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func makeTreeBranch(level int, last os.FileInfo, element os.FileInfo) string {
	var stringsBuilder strings.Builder
	for i := 0; i < level; i++ {
		fmt.Print("|")

		fmt.Print(&stringsBuilder, "\t")
	}
	if element == last {
		fmt.Print("└--")
	} else {
		fmt.Print("├--")
	}
	return stringsBuilder.String()
}

func showDir(out io.Writer, path string, printFiles bool, level int) error {
	dirContent, err := ioutil.ReadDir(path)

	if err != nil {
		return errors.New("can't read dir")
	}

	// remove files
	if !printFiles {
		for i := 0; i < len(dirContent); i++ {
			if !dirContent[i].IsDir() {
				dirContent = append(dirContent[:i], dirContent[i+1:]...)
				i--
			}
		}
	}


	for _, element := range dirContent {
		if element.IsDir() {
			fmt.Println(makeTreeBranch(level, dirContent[len(dirContent) - 1], element), element.Name())
			showDir(out, path + "/" + element.Name(), printFiles, level + 1)
		} else {
			fmt.Print(makeTreeBranch(level, dirContent[len(dirContent) - 1], element), element.Name())
			if element.Size() == 0 {
				fmt.Print(" (empty)\n")
			} else {
				fmt.Print(" (", element.Size(), "b)\n")
			}
		}
	}
	return nil
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	initialLevel := 0

	return showDir(out, path, printFiles, initialLevel)
}