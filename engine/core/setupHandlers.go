package core

import (
	"fmt"
	"net/http"

	"../../admin"
	"../../components"
	"../../modules"
	"../settings"
)

/*--------------------------
     HANDLERS SETUP FILE
  (What r u doing here-_-)
--------------------------*/

// InitHandlers makes all routes work(WOW!)
func InitHandlers() {

	routers := ParseRoutes()

	for _, value := range routers {
		template := fmt.Sprintf("./%s/%s", settings.TEMPLATES_DIR, value.Template)
		if IsTemplateExists(template) {
			http.HandleFunc(value.Adress, func(w http.ResponseWriter, r *http.Request) {
				http.ServeFile(w, r, template)
			})
		}
	}

	modules.SetCustomRoutes()

	if settings.ENABLE_ADMIN {
		http.HandleFunc(settings.ADMIN_ADRESS, admin.ServeAdminPage)
	}

	http.Handle(settings.STATIC_PREFIX, http.StripPrefix(settings.STATIC_PREFIX, http.FileServer(http.Dir(settings.STATIC_DIR))))

	components.PrintSuccess(fmt.Sprintf("Server is listening on http://127.0.0.1:%d", settings.PORT))
	fmt.Println("Stop with CTRL + C")
	http.ListenAndServe(fmt.Sprintf(":%d", settings.PORT), nil)

}
