package main

import (
	"fmt"
	"os"
)

func main() {
	gopath := os.Getenv("GOPATH")
	goroot := os.Getenv("GOROOT")

	fmt.Println(gopath)
	fmt.Println(goroot)
}

func GoExtList() {

}

func GoExtUnget() {

}
