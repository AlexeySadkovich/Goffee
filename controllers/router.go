package controllers

import (
	"../engine/core"
	"./handlers"
)

var route = core.Route

// SetCustomRoutes consists routes defined by user
func SetCustomRoutes() {
	/*---------------------------------
	 Put here your routes and handlers
	---------------------------------*/

	route("/", handlers.RenderIndexPage)
	route("/example", handlers.RenderExample)

}
