package test

import (
	"github.com/k0kubun/pp"
)

type user struct {
	name string
	age uint
}

func Status() {
	root := user{name: "Alex", age: 13}
	pp.Println(root)
}
