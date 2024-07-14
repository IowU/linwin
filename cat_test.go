package main

import "testing"

func TestCat(t *testing.T) {
	t.Run("Return contents of a single sample file", func(t *testing.T) {
		got := cat(`.\test_dir\test_file.txt`)
		want := "This is line 1\r\nThis is line 2\r\nThis is line 3\r\nThis is line 4\r\nThis is line 5\r\nThis is line 6\r\nThis is line 7\r\nThis is line 8\r\nThis is line 9\r\nThis is line 10\r\nThis is line 11"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Return contents of multiple files", func(t *testing.T) {
		got := cat(`.\test_dir\test_file.txt`, `.\test_dir\test_file_2.txt`)
		want := "This is line 1\r\nThis is line 2\r\nThis is line 3\r\nThis is line 4\r\nThis is line 5\r\nThis is line 6\r\nThis is line 7\r\nThis is line 8\r\nThis is line 9\r\nThis is line 10\r\nThis is line 11\r\nThis is file 2 line 1\r\nThis is file 2 line 2\r\nThis is file 2 line 3\r\nThis is file 2 line 4\r\nThis is file 2 line 5\r\nThis is file 2 line 6\r\nThis is file 2 line 7\r\nThis is file 2 line 8\r\nThis is file 2 line 9\r\nThis is file 2 line 10\r\nThis is file 2 line 11"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})
}
