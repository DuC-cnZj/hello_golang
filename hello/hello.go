package hello

import version "github.com/DuC-cnZj/hello_golang/vresion"

func Say() string {
	return "hello go" + version.GetVersion()
}

func run() string {
	return "run"
}