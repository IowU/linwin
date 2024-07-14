package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Prints the full contents of the specified files.
// Supports multiple files as input
func cat(path ...string) string {

	var result []string

	for _, item := range path {
		o, err := os.Open(item)

		if err != nil {
			log.Fatalf("ERROR: %v", err)
		}

		b, err := io.ReadAll(o)
		result = append(result, string(b))

	}

	return strings.Join(result, "\r\n")

}

func main() {
	fmt.Println(cat(`C:\Users\andre\Desktop\Projects\Go\linwin\test_dir\test_file.txt`))
}
