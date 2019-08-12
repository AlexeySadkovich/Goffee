package core

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"../../components"
	"../settings"
)

func ParseRoutes() []components.Route {

	if _, err := os.Stat(settings.ROUTES_PATH); os.IsNotExist(err) {
		components.PrintError("File with routes not found")
	}

	file, _ := ioutil.ReadFile(settings.ROUTES_PATH)

	routers := []components.Route{}

	_ = json.Unmarshal([]byte(file), &routers)

	return routers

}
