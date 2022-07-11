package main

import (
	"fmt"
	"io/ioutil"

	"git.sr.ht/~bossley9/sn/pkg/api"
)

func main() {
	fmt.Println("Hello, world!")

	params := map[string]string{}
	headers := map[string]string{}

	resp, err := api.Fetch("https://swapi.dev/api/people/1", "GET", params, headers)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}
