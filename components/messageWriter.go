package components

import (
	"log"

	"../engine/settings"
)

//PrintSuccess : Print OK messages
func PrintSuccess(message string) {
	log.Println("[  \033[92mOK\033[0m  ]", message)
}

//PrintInfo : Print INFO messages
func PrintInfo(message string) {
	if settings.DEBUG_MODE {
		log.Println("[ \033[94mINFO\033[0m ]", message)
	}
}

//PrintError : Print ERROR messages
func PrintError(message string) {
	if settings.DEBUG_MODE {
		log.Println("[ \033[91mERROR\033[0m ]", message)
	}
}
