package main

import (
	"fmt"
	"github.com/DuC-cnZj/hello_golang/version"
)

func main() {
	fmt.Println("hello: " + version.GetVersion())
}
