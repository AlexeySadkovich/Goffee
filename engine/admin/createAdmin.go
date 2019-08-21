package admin

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"

	"../../components"
	"../settings"
)

func CreateAdmin() {

	var answer string

	fmt.Println("New administrator")
	fmt.Println()

	db, err := sql.Open("sqlite3", settings.DB_PATH)
	if err != nil {
		components.PrintError(fmt.Sprintf("%s", err))
	}
	defer db.Close()

	username := getUsername()
	hashedPassword := GetMD5Hash(generatePassword())

	fmt.Print("Confirm?[yes]: ")
	fmt.Scan(&answer)

	if answer != "yes" && answer != "y" {
		fmt.Print("\033[91mAborted")
		return
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS 'Admin' ('Id' INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE, 'Username'	CHAR(30) UNIQUE, 'Password' CHAR(50))")
	if err != nil {
		components.PrintError(fmt.Sprintf("Unable to create table: %s", err))
		return
	}

	_, err = db.Exec("INSERT INTO Admin (username, password) values ($1, $2)", username, hashedPassword)
	if err != nil {
		components.PrintError(fmt.Sprintf("%s", err))
		return
	}

	fmt.Printf("\033[92mAdministrator %s successfully created!\033[0m", username)

}

func getUsername() string {

	var username string

	fmt.Print("Enter username(default: admin) > ")
	fmt.Scanln(&username)

	if len(username) == 0 {
		username = "admin"
	}

	return username

}

func generatePassword() string {

	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSUVWXYZ"

	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 7)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	password := string(b)

	fmt.Println("Your password:", password)

	return password

}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
