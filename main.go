package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func getModuleName() (string, error) {
	if len(os.Args) < 3 {
		return "", errors.New("please provide module name: witch [new|init] <module-name>")
	}
	return os.Args[2], nil
}

func createProjectStructure(moduleName string, createDir bool) error {
	if createDir {
		if err := os.MkdirAll(moduleName, 0755); err != nil {
			return fmt.Errorf("failed to create directory: %v", err)
		}
		if err := os.Chdir(moduleName); err != nil {
			return fmt.Errorf("failed to change directory: %v", err)
		}
	}

	if err := initGoMod(moduleName); err != nil {
		return err
	}

	if err := createMainFile(); err != nil {
		return err
	}

	fmt.Println("Project created successfully!")
	return nil
}

func initGoMod(moduleName string) error {
	cmd := exec.Command("go", "mod", "init", moduleName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run 'go mod init': %v", err)
	}
	return nil
}

func createMainFile() error {
	content := `package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}
`
	f, err := os.Create("main.go")
	if err != nil {
		return fmt.Errorf("failed to create main.go: %v", err)
	}
	defer f.Close()

	if _, err := f.WriteString(content); err != nil {
		return fmt.Errorf("failed to write to main.go: %v", err)
	}

	return nil
}

func cmdNew() error {
	moduleName, err := getModuleName()
	if err != nil {
		return err
	}
	return createProjectStructure(moduleName, true)
}

func cmdInit() error {
	moduleName, err := getModuleName()
	if err != nil {
		return err
	}
	return createProjectStructure(moduleName, false)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(`Это программа для создания проектов.
Использование: witch [КОМАНДА] [название проекта]

Список доступных команд:
  new <название проекта>   - создаёт директорию с проектом
  init <название проекта>  - создаёт проект в текущей директории`)
		return
	}

	var err error
	switch os.Args[1] {
	case "new":
		err = cmdNew()
	case "init":
		err = cmdInit()
	default:
		fmt.Println("Неизвестная команда:", os.Args[1])
		fmt.Println(`Доступные команды: new, init`)
		os.Exit(1)
	}

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
