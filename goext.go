package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	gopath := os.Getenv("GOPATH")
	goroot := os.Getenv("GOROOT")

	fmt.Println(gopath)
	fmt.Println(goroot)

	goPathRead := fmt.Sprintf("%s/src", gopath)

	fileList := []string{}
	err := filepath.Walk(goPathRead, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	})
	if err != nil {
		fmt.Println("err")
	}

	for _, file := range fileList {
		fmt.Println(file)
	}

}

func GoExtList() {

}

func GoExtUnget() {

}
