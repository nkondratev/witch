package utils

import (
	"errors"
	"os"
)

func getDirectoryName() (string, error) {
	if len(os.Args) < 3 {
		return "", errors.New("please provide directory name: witch [new|init] <directory-name>")
	}
	return os.Args[2], nil
}
