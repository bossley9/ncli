package main

import c "git.sr.ht/~bossley9/sn/pkg/client"

func main() {
	client := c.NewClient()
	client.DownloadSync()
}
