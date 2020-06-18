package hello

import (
	"github.com/DuC-cnZj/hello_golang/v2/version"
)

func Say() string {
	return "hello go" + version.GetVersion()
}

func run() string {
	return "run"
}