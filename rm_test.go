package main

import (
	"os"
	"testing"
)

func TestRm(t *testing.T) {
	t.Run("Attempts to remove a file from the single provided path", func(t *testing.T) {
		//First creates an empty file so to be removed later
		file, err := os.Create(`.\test_dir\to_be_removed_1.txt`)

		if err != nil {
			t.Errorf("Error while creating a temporary file: %q", err)
		}

		// If close is not called now, subsequent actions are prevented as
		// it is considered as used by another process
		file.Close()

		err_rm := rm(`.\test_dir\to_be_removed_1.txt`)

		if err_rm != nil {
			t.Errorf("Couldn't remove the file: %q", err_rm)
		}
	})

	t.Run("Attempts to remove files from multiple paths", func(t *testing.T) {
		//First creates empty files so to be removed later
		file, err := os.Create(`.\test_dir\to_be_removed_1.txt`)

		file2, err2 := os.Create(`.\test_dir\to_be_removed_2.txt`)

		if err != nil {
			t.Errorf("Error while creating a temporary file: %q", err)
		}

		if err2 != nil {
			t.Errorf("Error while creating a temporary file: %q", err2)
		}

		// If close is not called now, subsequent actions are prevented as
		// it is considered as used by another process
		file.Close()
		file2.Close()

		err_rm := rm(`.\test_dir\to_be_removed_1.txt`, `.\test_dir\to_be_removed_2.txt`)

		if err_rm != nil {
			t.Errorf("Couldn't remove the file: %q", err_rm)
		}

	})
}
