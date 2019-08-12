package core

import (
	"os"

	"../../components"
)

func IsTemplateExists(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		components.PrintError("Template " + file + " not found")
		return false
	} else {
		return true
	}
}
