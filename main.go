package main

import (
	//"io"
	"fmt"
	"os"
	//"path/filepath"
	"strings"
)

func dirTreeParser(out *os.File, dir string, counter int, printFiles bool) error {
	data, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	// Если printFiles == false, оставляем только каталоги
	var entries []os.DirEntry
	if !printFiles {
		for _, entry := range data {
			if entry.IsDir() {
				entries = append(entries, entry)
			}
		}
	} else {
		entries = data
	}

	for i, file := range entries {
		isLast := i == len(entries) - 1
		tab := strings.Repeat("|   ", counter)

		branch := "├───"
		if isLast {
			branch = "└───"
		}

		fmt.Fprintln(out, tab+branch+file.Name())

		// Рекурсивный вызов для каталогов
		if file.IsDir() {
			err := dirTreeParser(out, dir+"/"+file.Name(), counter+1, printFiles)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func dirTree(out *os.File, path string, printFiles bool) error {
	dirTreeParser(out, path, 0, printFiles)
	return nil
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
