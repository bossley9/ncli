package main

import (
	"fmt"
	"os"
	"strings"

	c "git.sr.ht/~bossley9/ncli/pkg/client"
)

func usage() {
	var output strings.Builder

	output.WriteString("usage:\n")
	output.WriteString("\tncli (command)\n")
	output.WriteString("commands:\n")
	output.WriteString("\t[no args]\tdownload and sync local files\n")

	fmt.Println(output.String())
}

func main() {
	client := c.NewClient()

	args := os.Args
	if len(args) <= 1 {
		err := client.DownloadSync()
		if err != nil {
			fmt.Println(err)
		}
	} else {
		usage()
	}
}
