package cmd

import (
	"bytes"
	"log"
	"testing"
	// "github.com/spf13/cobra"
)

func TestCat(t *testing.T) {

	type test struct {
		args []string
	}

	oneFile := test{
		args: []string{"cat", `..\test_dir\test_file.txt`},
	}

	t.Run("Return contents of a single sample file", func(t *testing.T) {
		cmd := rootCmd
		actual := new(bytes.Buffer)
		cmd.SetOut(actual)
		cmd.SetArgs(oneFile.args)
		cmd.Execute()

		got := actual.String()
		// 		expected :=
		// 			`This is line 1
		// This is line 2
		// This is line 3
		// This is line 4
		// This is line 5
		// This is line 6
		// This is line 7
		// This is line 8
		// This is line 9
		// This is line 10
		// This is line 11`

		expected := "This is line 1\nThis is line 2\nThis is line 3\nThis is line 4\nThis is line 5\nThis is line 6\nThis is line 7\nThis is line 8\nThis is line 9\nThis is line 10\nThis is line 11"

		if got != expected {
			log.Fatalf("got: %s, expected: %s", got, expected)
		}
	})
}
