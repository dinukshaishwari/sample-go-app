package errors

import (
	"log"
	"os"
)

func HandleIfError(errorToHandle error) {
	if errorToHandle == nil {
		return
	}
	log.Fatal(errorToHandle.Error())
	os.Exit(1)
}
