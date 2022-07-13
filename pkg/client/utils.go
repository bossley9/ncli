package client

import (
	"fmt"
	"os"
)

func getRootDir() string {
	home := os.Getenv("HOME")
	if len(home) == 0 {
		home = "."
	}
	return fmt.Sprintf("%s/Documents/notion", home)
}
