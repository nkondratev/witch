package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func createProjectStructure(moduleName string, directoryName string) error {
	if directoryName != "" {
		if err := os.MkdirAll(directoryName, 0755); err != nil {
			return fmt.Errorf("failed to create directory: %v", err)
		}
		if err := os.Chdir(directoryName); err != nil {
			return fmt.Errorf("failed to change directory: %v", err)
		}
	}

	if err := initGoMod(moduleName); err != nil {
		return err
	}

	if err := createMainFile(); err != nil {
		return err
	}

	if err := gitCommit(); err != nil {
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
	files := []string{"main.go", ".gitignore"}
	content := []string{`package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}
`, `# If you prefer the allow list template instead of the deny list, see community template:
# https://github.com/github/gitignore/blob/main/community/Golang/Go.AllowList.gitignore
#
# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with ` + "`go test -c`" +
		`
*.test
# Code coverage profiles and other test artifacts
*.out
coverage.*
*.coverprofile
profile.cov

# Dependency directories (remove the comment below to include it)
# vendor/

# Go workspace file
go.work
go.work.sum

# env file
.env

# Editor/IDE
# .idea/
# .vscode/
`}
	for i := range content {
		f, err := os.Create(files[i])
		if err != nil {
			return fmt.Errorf("failed to create %v: %v", files[i], err)
		}
		defer f.Close()

		if _, err := f.WriteString(content[i]); err != nil {
			return fmt.Errorf("failed to write to main.go: %v", err)
		}

	}
	return nil
}
