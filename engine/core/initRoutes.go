package core

import (
	"net/http"

	"../../components"
)

func Route(pattern string, handler func(http.ResponseWriter, *http.Request)) {

	components.Routes = append(components.Routes, pattern)

	http.HandleFunc(pattern, handler)

}
