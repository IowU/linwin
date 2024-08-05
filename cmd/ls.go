package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
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

		recursive, _ := cmd.Flags().GetBool("recursive")

		switch recursive {
		case true:
			if recursive == true {
				for _, path := range args {
					root := path
					err := filepath.WalkDir(root, visit)
					if err != nil {
						log.Fatalf("ERROR: %v", err)
					}
				}
			}
		case false:
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
		default:
			log.Fatalf("ERROR: recursive flag has an unexpected value %v", recursive)
		}
	},
}

//FLAGS definition

var r = cmdLs.Flags().BoolP("recursive", "r", false, "Enables or not a recursive listing of items under the specified path if it is a directory")

// WalkDirFunc defined here to reduce code clutter
func visit(path string, di fs.DirEntry, err error) error {
	if !strings.HasPrefix(path, ".") {

		fmt.Printf("%s\n", path)
	}
	return err
}
