package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"../../components"
	"../../engine/settings"
)

/*-------------------
  FILE FOR HANDLERS
-------------------*/

// RenderExample is example handler function
func RenderExample(w http.ResponseWriter, r *http.Request) {
	templateName := fmt.Sprintf("./%s/%s", settings.TEMPLATES_DIR, "example.html")
	tmpl, err := template.ParseFiles(templateName)
	if err != nil {
		components.PrintError(fmt.Sprintf("%s", err))
	}

	tmpl.Execute(w, "HELLO, Im example, you can delete me")
}
