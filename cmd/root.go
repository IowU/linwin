package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "linwin",
	Short: "linwin is my porting of some base linux commands to windows",
	Long: `Implementation is a work in progress. Don't use it in production environments
                or other critical environments.`,
}

func Execute() {
	rootCmd.Execute()

}

func init() {

	rootCmd.AddCommand(cmdCat)
	rootCmd.AddCommand(cmdRm)
	rootCmd.AddCommand(cmdLs)
}
