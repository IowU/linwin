package cmd

import (
	"bytes"
	"log"
	"testing"
)

func TestCat(t *testing.T) {

	type test struct {
		args []string
	}

	oneFile := test{
		args: []string{"cat", `..\test_dir\test_file.txt`},
	}

	multipleFiles := test{
		args: []string{"cat", `..\test_dir\test_file.txt`, `..\test_dir\test_file_2.txt`},
	}

	t.Run("Return contents of a single sample file", func(t *testing.T) {
		cmd := rootCmd
		actual := new(bytes.Buffer)
		cmd.SetOut(actual)
		cmd.SetArgs(oneFile.args)
		cmd.Execute()

		got := actual.String()

		expected := "This is line 1\r\nThis is line 2\r\nThis is line 3\r\nThis is line 4\r\nThis is line 5\r\nThis is line 6\r\nThis is line 7\r\nThis is line 8\r\nThis is line 9\r\nThis is line 10\r\nThis is line 11\n"

		if got != expected {
			// %q used instead of %v or %s as to enable the print of special characters
			log.Fatalf("got: %q, expected: %q", got, expected)
		}
	})

	t.Run("Return contents of multiple files", func(t *testing.T) {
		cmd := rootCmd
		actual := new(bytes.Buffer)
		cmd.SetOut(actual)
		cmd.SetArgs(multipleFiles.args)
		cmd.Execute()

		got := actual.String()

		expected := "This is line 1\r\nThis is line 2\r\nThis is line 3\r\nThis is line 4\r\nThis is line 5\r\nThis is line 6\r\nThis is line 7\r\nThis is line 8\r\nThis is line 9\r\nThis is line 10\r\nThis is line 11\r\nThis is file 2 line 1\r\nThis is file 2 line 2\r\nThis is file 2 line 3\r\nThis is file 2 line 4\r\nThis is file 2 line 5\r\nThis is file 2 line 6\r\nThis is file 2 line 7\r\nThis is file 2 line 8\r\nThis is file 2 line 9\r\nThis is file 2 line 10\r\nThis is file 2 line 11\n"

		// %q used instead of %v or %s as to enable the print of special characters
		if got != expected {
			log.Fatalf("got: %q, expected: %q", got, expected)
		}
	})

}
