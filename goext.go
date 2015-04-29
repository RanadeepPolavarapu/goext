package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	gopath := os.Getenv("GOPATH")

	if len(gopath) == 0 {
		fmt.Fprintf(os.Stdout, "Error: GOPATH is not set! Please set GOPATH environment variable and try again.\n")
		os.Exit(1)
	}
	// goroot := os.Getenv("GOROOT")

	goPathRead := fmt.Sprintf("%s/src", gopath)

	fileList := []string{}

	// The required walk Function for the filepath.Walk(...) function.
	walkFn := func(path string, info os.FileInfo, err error) error {
		fileList = AppendUniqueStringArray(fileList, filepath.Dir(path))
		return nil
	}
	if err := filepath.Walk(goPathRead, walkFn); err != nil {
		fmt.Println("err occurred in filepath Walk!")
	}

	for _, file := range fileList {
		fmt.Println(file)
	}

}

func GoExtList() {

}

func GoExtUnget() {

}

func AppendUniqueStringArray(slice []string, s string) []string {
	for _, element := range slice {
		if element == s {
			return slice
		}
	}
	return append(slice, s)
}
