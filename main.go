package main

import (
	"fmt"
	"os"

	"./components"
	"./controllers"
	"./engine/admin"
	"./engine/core"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	fmt.Println("|- Goffee")
	fmt.Println("|- Version: 0.1")
	fmt.Println("|")

	if len(os.Args) >= 2 {
		action := os.Args[1]

		if action == "create-admin" {
			admin.CreateAdmin()
		} else {
			components.PrintError("Unknown command")
		}
	} else {
		components.SetStartTime()
		controllers.SetCustomRoutes()
		core.InitHandlers()
	}

}
