package admin

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"../../components"
	"../settings"
	"github.com/gorilla/sessions"
)

/*----------------
 YOU ARE SOO DEEP
	  > <
       o
	be gentle!
----------------*/

type AdminInfo struct {
	Uptime      string
	Tables      []string
	Routes      []string
	AdminAdress string
}

type TableInfo struct {
	Name    string
	Columns []ColumnInfo
	Rows    [][]string
}

type ColumnInfo struct {
	Cid     int
	Name    string
	Type    string
	Notnull bool
	Default sql.NullString
	Pk      bool
}

type LoginData struct {
	Login    string
	Password string
}

var store = sessions.NewCookieStore([]byte("SESSION_KEY"))

// ControlAdmin decides where user will be redirected
func ControlAdmin(w http.ResponseWriter, r *http.Request) {

	session, err := store.Get(r, "MainSession")

	if err != nil {
		components.PrintError(fmt.Sprintf("%s", err))
		return
	}

	if auth, ok := session.Values["authenticated"].(bool); ok && auth {
		serveAdminPage(w, r)
	} else {
		serveLoginPage(w, r)
	}

}

/*------------
  ADMIN PAGE
------------*/

// ServeAdminPage send page which consists administator's tools
func serveAdminPage(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("./engine/admin/template/admin.html")
	if err != nil {
		components.PrintError(fmt.Sprintf("%s", err))
	}

	data := AdminInfo{
		Uptime:      components.GetUptime(),
		Tables:      getTables(),
		Routes:      components.Routes,
		AdminAdress: settings.ADMIN_ADRESS,
	}

	tmpl.Execute(w, data)

}

/*-------
  LOGIN
-------*/

func serveLoginPage(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("./engine/admin/template/login.html")
	if err != nil {
		components.PrintError(fmt.Sprintf("%s", err))
	}

	data := AdminInfo{
		AdminAdress: settings.ADMIN_ADRESS,
	}

	tmpl.Execute(w, data)

}

// LoginHandler checks auth data
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		var rightPassword string
		var data LoginData

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			components.PrintError(fmt.Sprintf("%s", err))
			return
		}

		_ = json.Unmarshal(body, &data)

		username := data.Login
		password := data.Password

		session, err := store.Get(r, "MainSession")
		if err != nil {
			components.PrintError(fmt.Sprintf("Session error :%s", err))
			return
		}

		db, err := sql.Open("sqlite3", settings.DB_PATH)
		components.CheckError(err)
		defer db.Close()

		query := fmt.Sprintf("SELECT Password FROM Admin WHERE Username = '%s';", username)
		res, err := db.Query(query)
		if err != nil {
			components.PrintError(fmt.Sprintf("Quering failed: %s", err))
		}
		defer res.Close()

		for res.Next() {
			err := res.Scan(&rightPassword)
			if err != nil {
				components.PrintError(fmt.Sprintf("%s", err))
				continue
			}
		}

		if GetMD5Hash(password) == rightPassword {
			session.Values["username"] = username
			session.Values["authenticated"] = true
			session.Save(r, w)

			response, err := json.Marshal([]int{200})
			if err != nil {
				components.PrintError(fmt.Sprintf("%s", err))
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(response)
		}

	}
}

func ChangeAdminPassword(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		var rightPassword string
		var data struct {
			OldPassword string
			NewPassword string
		}

		session, err := store.Get(r, "MainSession")
		if err != nil {
			components.PrintError(fmt.Sprintf("Session error :%s", err))
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			components.PrintError(fmt.Sprintf("%s", err))
			return
		}

		_ = json.Unmarshal(body, &data)

		username := session.Values["username"]
		oldPassword := data.OldPassword
		newPassword := data.NewPassword

		db, err := sql.Open("sqlite3", settings.DB_PATH)
		components.CheckError(err)
		defer db.Close()

		query := fmt.Sprintf("SELECT Password FROM Admin WHERE Username = '%s';", username)
		res, err := db.Query(query)
		if err != nil {
			components.PrintError(fmt.Sprintf("Quering failed: %s", err))
		}
		defer res.Close()

		for res.Next() {
			err := res.Scan(&rightPassword)
			if err != nil {
				components.PrintError(fmt.Sprintf("%s", err))
				continue
			}
		}

		if GetMD5Hash(oldPassword) == rightPassword {

			query := fmt.Sprintf("UPDATE Admin SET Password = '%s' WHERE Username = '%s'", GetMD5Hash(newPassword), username)
			_, err := db.Exec(query)
			if err != nil {
				components.PrintError(fmt.Sprintf("%s", err))
				return
			}

			response, err := json.Marshal([]int{200})
			if err != nil {
				components.PrintError(fmt.Sprintf("%s", err))
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(response)

		}

	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	session, err := store.Get(r, "MainSession")
	if err != nil {
		components.PrintError(fmt.Sprintf("Session error :%s", err))
		return
	}

	session.Values["authenticated"] = false
	session.Save(r, w)

	response, err := json.Marshal([]int{200})
	if err != nil {
		components.PrintError(fmt.Sprintf("%s", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}

/*--------
  TABLES
--------*/

// ServeTableInfoPage generates and sends page with table strucure
func ServeTableInfoPage(w http.ResponseWriter, r *http.Request) {

	tableName := r.URL.Query().Get("table")

	tmpl, err := template.ParseFiles("./engine/admin/template/table.html")
	if err != nil {
		components.PrintError(fmt.Sprintf("%s", err))
	}

	data := TableInfo{
		Name:    tableName,
		Columns: getTableColumns(tableName),
		Rows:    getTableRows(tableName),
	}

	tmpl.Execute(w, data)

}

func DeleteTable(w http.ResponseWriter, r *http.Request) {

	session, err := store.Get(r, "MainSession")
	if err != nil {
		components.PrintError(fmt.Sprintf("Session error :%s", err))
		return
	}

	if auth, ok := session.Values["authenticated"].(bool); ok && auth {
		var data struct {
			Table string
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			components.PrintError(fmt.Sprintf("%s", err))
			return
		}

		_ = json.Unmarshal(body, &data)

		table := data.Table

		db, err := sql.Open("sqlite3", settings.DB_PATH)
		if err != nil {
			components.PrintError(fmt.Sprintf("Error while deleting table: %s", err))
			return
		}
		defer db.Close()

		query := fmt.Sprintf("DROP TABLE %s", table)
		_, err = db.Exec(query)
		if err != nil {
			components.PrintError(fmt.Sprintf("Error while deleting table: %s", err))
			return
		}

		response, err := json.Marshal([]int{200})
		if err != nil {
			components.PrintError(fmt.Sprintf("%s", err))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}
}

func getTableRows(table string) [][]string {

	var rows = make([][]string, getAmountOfRows(table))

	cols := getTableColumns(table)

	db, err := sql.Open("sqlite3", settings.DB_PATH)
	components.CheckError(err)
	defer db.Close()

	for _, value := range cols {
		query := fmt.Sprintf("SELECT %s FROM %s", value.Name, table)
		res, err := db.Query(query)
		components.CheckError(err)
		defer res.Close()

		r := 0

		for res.Next() {
			var field []byte
			err := res.Scan(&field)
			if err != nil {
				components.PrintError(fmt.Sprintf("%s", err))
				continue
			}
			rows[r] = append(rows[r], string(field))
			r++
		}
	}

	return rows

}

func getTableColumns(table string) []ColumnInfo {

	var columns []ColumnInfo

	db, err := sql.Open("sqlite3", settings.DB_PATH)
	components.CheckError(err)
	defer db.Close()

	query := fmt.Sprintf("PRAGMA table_info(%s)", table)
	res, err := db.Query(query)
	components.CheckError(err)
	defer res.Close()

	for res.Next() {
		c := ColumnInfo{}
		err := res.Scan(&c.Cid, &c.Name, &c.Type, &c.Notnull, &c.Default, &c.Pk)
		if err != nil {
			components.PrintError(fmt.Sprintf("%s", err))
			continue
		}
		columns = append(columns, c)
	}

	return columns

}

func getAmountOfRows(table string) int {

	db, err := sql.Open("sqlite3", settings.DB_PATH)
	components.CheckError(err)
	defer db.Close()

	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", table)

	res, err := db.Query(query)
	components.CheckError(err)
	defer res.Close()

	amount := 0

	for res.Next() {
		err := res.Scan(&amount)
		if err != nil {
			components.PrintError(fmt.Sprintf("%s", err))
			continue
		}
	}

	return amount

}

func getTables() []string {

	var tables []string

	db, err := sql.Open("sqlite3", settings.DB_PATH)
	components.CheckError(err)
	defer db.Close()

	res, err := db.Query("SELECT name FROM sqlite_master WHERE type ='table' AND name NOT LIKE 'sqlite_%'")
	components.CheckError(err)
	defer res.Close()

	for res.Next() {
		tableName := ""
		err := res.Scan(&tableName)
		if err != nil {
			components.PrintError(fmt.Sprintf("%s", err))
			continue
		}

		tables = append(tables, tableName)
	}

	return tables

}
