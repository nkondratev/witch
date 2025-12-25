/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/nkondratev/witch/utils"
	"github.com/spf13/cobra"
)

// libCmd represents the lib command
var libCmd = &cobra.Command{
	Use:   "lib",
	Short: "Use to create directory with your github user name in go mod for library",
	Long: `Use to create directory with your github user name in go mod for library
Examples:
witch lib images
witch lib driver
`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.CmdLib()
	},
}

func init() {
	rootCmd.AddCommand(libCmd)
}
