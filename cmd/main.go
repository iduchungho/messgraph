package main

import (
	"messgraph.com/m/internal/engine"
)

func main() {
	engine := engine.NewEngine()
	engine.Prepare()
	engine.Run()
}