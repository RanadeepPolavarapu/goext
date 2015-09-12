package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	EnvVarGOPATH = "GOPATH"
	EnvVarGOROOT = "GOROOT"
)

func main() {
	gopath := os.Getenv(EnvVarGOPATH)

	if len(gopath) == 0 {
		fmt.Fprintf(os.Stderr, "Error: GOPATH is not set! Please set GOPATH environment variable and try again.\n")
		os.Exit(1)
	}

	// goroot := os.Getenv(EnvVarGOROOT)

	goPathRead := fmt.Sprintf("%s/src", gopath)

	fileList := []string{}

	// The required walk Function for the filepath.Walk(...) function.
	walkFn := func(path string, info os.FileInfo, err error) error {
		dirOnlyPath := filepath.Dir(path)
		fileList = AppendUniqueStringArray(fileList, dirOnlyPath)
		return nil
	}
	if err := filepath.Walk(goPathRead, walkFn); err != nil {
		fmt.Fprintf(os.Stderr, "Error: err occurred in filepath Walk!\n")
	}

	for _, file := range fileList {
		fmt.Println(file)
	}

}

func GoExtList() {

}

func GoExtUnget() {

}

// Utility function to append only unique values into a string array.
func AppendUniqueStringArray(slice []string, s string) []string {
	for _, element := range slice {
		if element == s {
			return slice
		}
	}
	return append(slice, s)
}
