package main

import (
	//"fmt"
	//"io"
	"os"
	//"path/filepath"
	"strings"
)

func dirTreeTest(dir string, counter int) {
	data, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	for _, file := range data {
		
		if file.IsDir() {
			tab := strings.Repeat("|	", counter)
			os.Stdout.WriteString(tab + "├───" + file.Name() + "\n")
			counter++
			dirTreeTest(dir + "/" + file.Name(), counter)
		} else {
			tab := strings.Repeat("|	", counter)
			os.Stdout.WriteString(tab + "└───" + file.Name() + "\n")
		}
	}
}
govori!!!
func main() {
	// out := os.Stdout
	// if !(len(os.Args) == 2 || len(os.Args) == 3) {
	// 	panic("usage go run main.go . [-f]")
	// }
	// path := os.Args[1]
	// printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	// err := dirTree(out, path, printFiles)
	// if err != nil {
	// 	panic(err.Error())
	// }
	//result := &[]string{}
	dirTreeTest("./testdata", 0)
}
