package admin

import (
	"fmt"
	"html/template"
	"net/http"

	"../components"
)

type AdminInfo struct {
	Uptime string
}

func ServeAdminPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./admin/template/admin.html")
	if err != nil {
		components.PrintError(fmt.Sprintf("%s", err))
	}

	data := AdminInfo{
		Uptime: components.GetUptime(),
	}

	tmpl.Execute(w, data)
}
