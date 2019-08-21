package components

import "fmt"

func CheckError(err error) {
	if err != nil {
		PrintError(fmt.Sprintf("%s", err))
	}
}
