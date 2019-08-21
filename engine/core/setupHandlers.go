package core

import (
	"fmt"
	"net/http"

	"../../components"
	"../admin"
	"../settings"
)

/*--------------------------
     HANDLERS SETUP FILE
  (What r u doing here-_-)
--------------------------*/

// InitHandlers makes all routes work(WOW!)
func InitHandlers() {

	if settings.DEBUG_MODE {
		fmt.Println("|-[!]-Debug mode is \033[93menabled\033[0m. All errors will be shown.")
		fmt.Println()
	}

	components.PrintInfo("Initializing routes and handlers")

	if settings.ENABLE_ADMIN {
		http.HandleFunc(settings.ADMIN_ADRESS, admin.ControlAdmin)
		http.HandleFunc(settings.ADMIN_ADRESS+"/table", admin.ServeTableInfoPage)
		http.HandleFunc(settings.ADMIN_ADRESS+"/change-password", admin.ChangeAdminPassword)
		http.HandleFunc(settings.ADMIN_ADRESS+"/auth", admin.LoginHandler)
		http.HandleFunc(settings.ADMIN_ADRESS+"/deauth", admin.LogoutHandler)
		http.HandleFunc(settings.ADMIN_ADRESS+"/table/delete", admin.DeleteTable)
	}

	http.Handle(settings.STATIC_PREFIX, http.StripPrefix(settings.STATIC_PREFIX, http.FileServer(http.Dir(settings.STATIC_DIR))))

	components.PrintSuccess(fmt.Sprintf("Server is listening on http://127.0.0.1:%d", settings.PORT))
	http.ListenAndServe(fmt.Sprintf(":%d", settings.PORT), nil)

}
