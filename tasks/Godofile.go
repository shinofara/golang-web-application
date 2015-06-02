package main

import (
	. "gopkg.in/godo.v1"
)

func tasks(p *Project) {
	Env = `GOPATH=.vendor::$GOPATH`

	p.Task("server", D{}, func() {
		Start("main.go", M{"$in": "./"})
	}).Watch("*.go", "**/*.go").
		Debounce(3000)
}

func main() {
	Godo(tasks)
}
