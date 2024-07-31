package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
	"strings"
)

// Prints the full contents of the specified files.
// Supports multiple files as input
// TODO: return errors when an error occurs and handle them
// in the main function

var cmdCat = &cobra.Command{
	Use:   "cat [list of paths to files]",
	Short: "Prints the contents of the files passed as arguments",
	Long:  `cat is for printing the files contents to the screen, each of which concatenated with the previous one.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		var result []string

		for _, item := range args {
			o, err := os.Open(item)

			if err != nil {
				log.Fatalf("ERROR: %v", err)
			}

			b, err := io.ReadAll(o)
			result = append(result, string(b))

			if err != nil {
				log.Fatalf("ERROR: %v", err)
			}
		}
		fmt.Fprintln(cmd.OutOrStdout(), strings.Join(result, ""))
		// fmt.Println(strings.Join(result, "\r\n"))
	},
}
