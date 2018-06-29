package main

import (
	"os"
	"plugin"
)

func main() {
	gopath := os.Getenv("GOPATH")
	p, err := plugin.Open(gopath + "/src/go-study/bin/goplugin/goplugin.so")
	if err != nil {
		panic(err)
	}

	v, err := p.Lookup("V")
	if err != nil {
		panic(err)
	}

	f, err := p.Lookup("F")
	if err != nil {
		panic(err)
	}

	*v.(*int) = 87
	f.(func())()
}
