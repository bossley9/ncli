package main

import (
	"fmt"
	"os"
	"strings"

	c "git.sr.ht/~bossley9/sn/pkg/client"
)

func usage() {
	var output strings.Builder

	output.WriteString("usage:\n")
	output.WriteString("\tsn (command)\n")
	output.WriteString("commands:\n")
	output.WriteString("\t[no args]\tdownload and sync local files\n")

	fmt.Println(output.String())
}

func main() {
	client := c.NewClient()

	args := os.Args
	if len(args) <= 1 {
		client.DownloadSync()
	} else {
		usage()
	}
}
