package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// Prints the specified directory's content.
// Supports multiple directories.
// TODO: return errors when an error occurs and handle them
// in the main function
// TODO: implement flags as -l, -r

var cmdLs = &cobra.Command{
	Use:   "ls [list of paths to folders]",
	Short: "Prints the specified directory's content.",
	Long:  `ls prints the specified directory's content.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		for _, path := range args {
			item, err := os.Stat(path)

			if err != nil {
				log.Fatalf("ERROR: %v", err)
			}
			if item.IsDir() {

				element, err := os.ReadDir(path)
				if err != nil {
					log.Fatalf("ERROR: %v", err)

				}

				for _, entry := range element {
					fmt.Println(entry.Name())
				}
			}
		}
	},
}
