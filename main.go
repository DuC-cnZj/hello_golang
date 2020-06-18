package main

import (
	"fmt"
	version "github.com/DuC-cnZj/hello_golang/vresion"
)

func main() {
	fmt.Println("hello" + version.GetVersion())
}
