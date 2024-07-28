package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

// Prints the full contents of the specified files.
// Supports multiple files as input
// TODO: return errors when an error occurs and handle them
// in the main function
// TODO: implement flags as -r, -f

var cmdRm = &cobra.Command{
	Use:   "rm [list of paths to files or folders]",
	Short: "removes the specified files or folders",
	Long:  `rm removes the specified files or folders`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		recursive, _ := cmd.Flags().GetBool("recursive")

		for _, item := range args {

			if recursive == true {
				err := os.RemoveAll(item)

				if err != nil {
					log.Fatalf("ERROR: %v", err)
				}

			} else {

				err := os.Remove(item)

				if err != nil {
					log.Fatalf("ERROR: %v", err)
				}
			}
		}

	},
}

// func BoolP(name, shorthand string, value bool, usage string) *bool

var f = cmdRm.Flags().BoolP("recursive", "r", false, "Enables or not a recursive removal of all items under the specified path if it is a directory")
