package modules

import (
	"net/http"

	"./handlers"
)

// SetCustomRoutes consists routes defiened by user
func SetCustomRoutes() {
	/*---------------------------------
	 Put here your routes and handlers
	---------------------------------*/

	http.HandleFunc("/example", handlers.RenderExample)

}
