// The test should:
// Create a file in the test directory
// Attempt to remove the file created previously with the command defined in this application:
//   - if successful, check the file doesn't actually exist anymore
//   - if not successful, remove the file through something reliable as the os package
//
// repeat the above, but with multiple files specified as arguments
// test the flags

package cmd

import (
	"log"
	"os"
	"testing"
)

func TestRm(t *testing.T) {

	type test struct {
		args []string
	}

	oneFile := test{
		args: []string{"rm", `..\test_dir\removable_test_file.txt`},
	}

	// multipleFiles := test{
	// 	args: []string{"rm", `..\test_dir\removable_\test_file.txt`, `..\test_dir\removable_test_file_2.txt`},
	// }

	t.Run("Removes a single file", func(t *testing.T) {
		file, err := os.Create(oneFile.args[1])

		//Close the file or it won't be available for later interaction
		file.Close()

		if err != nil {
			log.Fatalf("ERROR: %s", err)
		}

		cmd := rootCmd
		cmd.SetArgs(oneFile.args)
		cmd.Execute()

		if _, err := os.Stat(oneFile.args[1]); err == nil {
			log.Fatalf("ERROR: the file %s stil exists", oneFile.args[1])
		}

	})

	// t.Run("Removes multiple file", func(t *testing.T) {
	// 	cmd := rootCmd
	// 	cmd.SetArgs(multipleFiles.args)
	// 	cmd.Execute()
	// })
}
