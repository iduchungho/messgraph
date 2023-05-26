package main

import (
	"messgraph.com/m/modules/engine"
)

func main() {
	app := engine.NewEngine()
	app.Prepare()
	app.Run()
}
