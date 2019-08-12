package main

import (
	"./components"
	"./engine/core"
)

func main() {
	components.PrintInfo("Initializing routes and handlers")
	components.SetStartTime()
	core.InitHandlers()
}
