package utils

func CmdNew() error {
	moduleName, err := getDirectoryName()
	if err != nil {
		return err
	}
	return createProjectStructure(moduleName, moduleName)
}
