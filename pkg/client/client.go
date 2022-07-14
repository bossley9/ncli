package client

import (
	"fmt"
	"os"
)

type Client struct {
	ProjectDir string
	Metadata   Metadata
}

func NewClient() *Client {
	fmt.Println("initializing client...")
	client := Client{}

	// set and create project directory
	home := os.Getenv("HOME")
	if len(home) == 0 {
		home = "."
	}
	client.ProjectDir = fmt.Sprintf("%s/Documents/notion", home)
	err := os.MkdirAll(client.ProjectDir, 0700)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// set metadata
	client.Metadata = Metadata{}

	fmt.Println("done.")
	return &client
}
