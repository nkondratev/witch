package utils

func CmdInit() error {
	moduleName, err := getDirectoryName()
	if err != nil {
		return err
	}
	return createProjectStructure(moduleName, "")
}
