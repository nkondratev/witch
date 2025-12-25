package utils

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

func CmdLib() error {
	directoryName, err := getDirectoryName()
	var name string
	if err != nil {
		panic("Cannot get module name " + err.Error())
	}
	home, err := os.UserHomeDir()
	if err != nil {
		panic("Cannot get home dir " + err.Error())

	}
	path := path.Join(home, "/.gitconfig")

	gitConfig, err := os.Open(path)
	if err != nil {
		panic("Cannot read gitconfig " + err.Error())
	}
	defer gitConfig.Close()
	scanner := bufio.NewScanner(gitConfig)
	var hasUser bool

	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(text, "[user]") {
			hasUser = true
		}
		if strings.HasPrefix(text, "name") && hasUser {
			parts := strings.SplitN(text, "=", 2)
			name = strings.TrimSpace(parts[1])
			fmt.Println(name)
		}
	}
	if name == "" {
		panic("error, your not set user name in /home/<user>/.gitconfig")
	}
	moduleName := "github.com/" + name + "/" + directoryName
	return createProjectStructure(moduleName, directoryName)
}
