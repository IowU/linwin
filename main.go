package main

import (
	// "fmt"
	"linwin/cmd"
	// "io"
	// "log"
	// "os"
	// "strings"
)

// TODO: return errors when an error occurs and handle them
// in the main function
// TODO: implement flags as -r, -f etc.
// func rm(path ...string) error {

// 	var err error
// 	for _, file := range path {

// 		err = os.Remove(file)

// 		if err != nil {
// 			log.Fatalf("ERROR: %q", err)
// 			break
// 		}

// 	}

// 	return err
// }

func main() {
	cmd.Execute()
}
