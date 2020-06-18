package main

import (
	"fmt"
	"github.com/DuC-cnZj/hello_golang/v2/version"
)

func main() {
	fmt.Println("hello: " + version.GetVersion())
}
