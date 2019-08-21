@echo off
echo Welcome to Goffee setup!

set /p id="Project name: "

ren "./main.go" "%id%.go"

echo [INFO] Installing dependencies...
cmd /c "go get github.com/mattn/go-sqlite3"
cmd /c "go get github.com/gorilla/sessions"
echo Done

echo [INFO] Compiling...
cmd /c "go build %id%.go"
echo Done

PAUSE